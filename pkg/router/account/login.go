package account

import (
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/model/account"
	model "ceres/pkg/model/account"
	"ceres/pkg/router"
	service "ceres/pkg/service/account"
	"ceres/pkg/utility/auth"
	"ceres/pkg/utility/jwt"
	"errors"
	"fmt"
	"strings"

	"github.com/qiniu/x/log"
)

// LoginWithGithubCallback login with github oauth
func LoginWithGithubCallback(ctx *router.Context) {
	loginWithOauth(ctx, func(code string) (auth.OauthAccount, error) {
		client := auth.NewGithubOauthClient(code)
		return client.GetUserProfile()
	})
}
func extractComerIdFromJwtToken(ctx *router.Context) (comerID uint64, err error) {
	comunioAuthHeader := ctx.GetHeader("X-COMUNION-AUTHORIZATION")
	if strings.Trim(comunioAuthHeader, " ") == "" {
		comerID = 0
		err = nil
	} else {
		comerID, err = jwt.Verify(comunioAuthHeader)
	}
	log.Debugf("EXTRACT COMUNION_COMER_ID FROM JWT_TOKEN: %d\n", comerID)
	return
}

// LoginWithGoogleCallback login with google oauth callback
func LoginWithGoogleCallback(ctx *router.Context) {
	loginWithOauth(ctx, func(code string) (auth.OauthAccount, error) {
		client := auth.NewGoogleClient(code)
		return client.GetUserProfile()
	})
}

// GetBlockchainLoginNonce get the blockchain login nonce.
func GetBlockchainLoginNonce(ctx *router.Context) {
	address := ctx.Query("address")
	if address == "" {
		err := router.ErrBadRequest.WithMsg("Invalid address")
		ctx.HandleError(err)
		return
	}

	var nonce account.WalletNonceResponse
	if err := service.GenerateWeb3LoginNonce(address, &nonce); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(nonce)
}

// LoginWithWallet login with the wallet signature.
func LoginWithWallet(ctx *router.Context) {
	var request model.EthLoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Warn(err)
		err = router.ErrBadRequest.WithMsg(err.Error())
		ctx.HandleError(err)
		return
	}

	var response account.ComerLoginResponse
	if err := service.LoginWithEthWallet(request.Address, request.Signature, &response); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(response)
}

// RegisterWithOauth 基于oauth帐号注册，创建comer以及profile
func RegisterWithOauth(ctx *router.Context) {
	var request model.RegisterWithOauthRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Warn(err)
		err = router.ErrBadRequest.WithMsg(err.Error())
		ctx.HandleError(err)
		return
	}
	var (
		comer         model.Comer
		comerAccount  model.ComerAccount
		comerProfile  model.ComerProfile
		loginResponse model.OauthLoginResponse
	)

	if err := account.GetComerAccountById(mysql.DB, request.OauthAccountId, &comerAccount); err != nil {
		handleError(ctx, err)
		return
	}
	if comerAccount.ID == 0 {
		handleError(ctx, errors.New("oauth account does not exist or has been deleted"))
		return
	}
	if comerAccount.ComerID != 0 {
		handleError(ctx, errors.New("account has bind to another comerId"))
		return
	}

	comer = model.Comer{
		Address: nil,
	}

	if err := account.CreateComer(mysql.DB, &comer); err != nil {
		handleError(ctx, err)
		return
	}
	if err := service.CreateComerProfile(comer.ID, &request.Profile); err != nil {
		handleError(ctx, err)
		return
	}
	comerAccount.ComerID = comer.ID
	comerAccount.IsLinked = true

	if err := service.LinkOauthToComer(comerAccount.ID, comer.ID); err != nil {
		handleError(ctx, err)
		return
	}

	if err := account.GetComerProfile(mysql.DB, comer.ID, &comerProfile); err != nil {
		handleError(ctx, err)
		return
	}

	loginResponse = model.OauthLoginResponse{
		ComerID:        comer.ID,
		Nick:           comerProfile.Name,
		Avatar:         comerProfile.Avatar,
		Address:        *comer.Address,
		Token:          jwt.Sign(comer.ID),
		IsProfiled:     true,
		OauthLinked:    true,
		OauthAccountId: comerAccount.ID,
	}
	ctx.OK(loginResponse)
	return
}

func handleError(ctx *router.Context, err error) {
	log.Warn(err)
	ctx.HandleError(router.ErrBadRequest.WithMsg(err.Error()))
	return
}

func loginWithOauth(ctx *router.Context, oauthAccount func(string) (auth.OauthAccount, error)) {
	code := ctx.Query("code")
	if code == "" {
		handleError(ctx, router.ErrBadRequest.WithMsg("Code missed"))
		return
	}
	if oauth, err := oauthAccount(code); err != nil {
		log.Warn(err)
		ctx.HandleError(router.ErrBadRequest.WithMsg("Login authorization failed"))
		return
	} else {
		var (
			comerAccount  model.ComerAccount
			comer         model.Comer
			comerProfile  model.ComerProfile
			loginResponse account.OauthLoginResponse
			comerId       uint64
		)
		if err = account.GetComerAccount(mysql.DB, model.GithubOauth, oauth.GetUserID(), &comerAccount); err != nil {
			handleError(ctx, err)
			return
		}
		loginResponse = account.OauthLoginResponse{}
		loginResponse.ComerID = 0
		loginResponse.IsProfiled = false
		loginResponse.OauthLinked = false
		loginResponse.OauthAccountId = comerAccount.ID
		// 首次登录
		if comerAccount.ID == 0 {
			comerAccount = model.ComerAccount{
				ComerID:   0,
				OIN:       oauth.GetUserID(),
				IsPrimary: true,
				Nick:      oauth.GetUserNick(),
				Avatar:    oauth.GetUserAvatar(),
				Type:      model.GithubOauth,
				IsLinked:  false,
			}
			if err := account.CreateAccount(mysql.DB, &comerAccount); err != nil {
				handleError(ctx, err)
				return
			}
		} else {
			// 2 已关联到comer
			if comerAccount.ComerID != 0 && comerAccount.IsLinked {
				if err := account.GetComerByID(mysql.DB, comerAccount.ComerID, &comer); err != nil {
					handleError(ctx, err)
					return
				}
				if comer.ID == 0 || comer.IsDeleted {
					handleError(ctx, errors.New(fmt.Sprintf("Comer does not exist or was deleted!")))
					return
				}
				if err := account.GetComerProfile(mysql.DB, comerAccount.ComerID, &comerProfile); err != nil {
					handleError(ctx, err)
					return
				}
				loginResponse.Nick = comerProfile.Name
				loginResponse.Avatar = comerProfile.Avatar
				loginResponse.IsProfiled = true
				loginResponse.ComerID = comer.ID
				loginResponse.Address = *comer.Address
				loginResponse.OauthLinked = true
				loginResponse.OauthAccountId = comerAccount.ID
			}
		}
		comerId = comerAccount.ComerID
		token := jwt.Sign(comerId)
		loginResponse.Token = token

		ctx.OK(loginResponse)
	}
}
