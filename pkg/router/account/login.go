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
	"gorm.io/gorm"
	"strings"

	"github.com/qiniu/x/log"
)

// LoginWithGithubCallback login with github oauth
func LoginWithGithubCallback(ctx *router.Context) {
	loginWithOauth(ctx, model.GithubOauth, func(code string) (auth.OauthAccount, error) {
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
	loginWithOauth(ctx, model.GoogleOauth, func(code string) (auth.OauthAccount, error) {
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

// RegisterWithOauth 基于oauth帐号注册，创建profile, 首次用oauth登录并不连接到已有Comer时候点取消时候的注册接口
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

	if err := account.GetComerByID(mysql.DB, comerAccount.ID, &comer); err != nil {
		handleError(ctx, err)
		return
	}
	if err := service.CreateComerProfile(comer.ID, &request.Profile); err != nil {
		handleError(ctx, err)
		return
	}
	comerAccount.ComerID = comer.ID
	comerAccount.IsLinked = true

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

func loginWithOauth(ctx *router.Context, oauthType model.ComerAccountType, oauthAccount func(string) (auth.OauthAccount, error)) {
	code := ctx.Query("code")
	if code == "" {
		handleError(ctx, router.ErrBadRequest.WithMsg("Code missed"))
		return
	}
	var (
		logonComerId uint64
		err          error
		oauth        auth.OauthAccount
	)
	if logonComerId, err = extractComerIdFromJwtToken(ctx); err != nil {
		logonComerId = 0
	}
	oauth, err = oauthAccount(code)

	if err != nil {
		log.Warn(err)
		ctx.HandleError(router.ErrBadRequest.WithMsg("Login authorization failed"))
		return
	}
	log.Debugf("loginWithOauth oauthInfo: %v\n", oauth)
	if logonComerId == 0 {
		err, loginResponse := loginWithUnRegistredComer(oauth, oauthType)
		if err != nil {
			handleError(ctx, err)
			return
		}
		ctx.OK(loginResponse)
		log.Debugf("loginWithOauth response: %v\n", loginResponse)
		return
	} else {
		err, loginResponse := loginWithRegisteredComer(oauth, oauthType, logonComerId)
		if err != nil {
			handleError(ctx, err)
			return
		}
		ctx.OK(loginResponse)
		log.Debugf("loginWithOauth when logon response: %v\n", loginResponse)
		return
	}

}

func loginWithRegisteredComer(oauth auth.OauthAccount, oauthType model.ComerAccountType, logonComerId uint64) (err error, loginResponse model.OauthLoginResponse) {

	var (
		comerAccount model.ComerAccount
		comer        model.Comer
		comerProfile model.ComerProfile
	)
	// link oauth after logon by wallet!!
	if err = model.GetComerByID(mysql.DB, logonComerId, &comer); err != nil {
		return
	}
	if comer.ID == 0 {
		err = errors.New(fmt.Sprintf("comer with id %d does not exist", logonComerId))
		return
	}
	var (
		linkedAccounts              []model.ComerAccount
		comerNotLinkedThisTypeOauth = true
	)
	if err = model.GetComerAccountsByComerId(mysql.DB, logonComerId, &linkedAccounts); err != nil {
		return
	}
	if linkedAccounts == nil || len(linkedAccounts) == 0 {
		comerNotLinkedThisTypeOauth = true
	} else {
		for _, linkedAccount := range linkedAccounts {
			if linkedAccount.Type == oauthType {
				comerNotLinkedThisTypeOauth = false
				break
			}
		}
	}
	loginResponse = model.OauthLoginResponse{
		ComerID:        logonComerId,
		Nick:           oauth.GetUserNick(),
		Avatar:         oauth.GetUserAvatar(),
		Address:        *comer.Address,
		Token:          jwt.Sign(logonComerId),
		IsProfiled:     false,
		OauthLinked:    false,
		OauthAccountId: comerAccount.ID,
	}
	if err = account.GetComerProfile(mysql.DB, logonComerId, &comerProfile); err != nil {
		return
	}
	if comerProfile.ID != 0 {
		loginResponse.IsProfiled = true
		loginResponse.Nick = comerProfile.Name
		loginResponse.Avatar = comerProfile.Avatar
	}
	if comerNotLinkedThisTypeOauth {
		// 检查当前oauthAccount能否被关联
		if err = model.GetComerAccount(mysql.DB, oauthType, oauth.GetUserID(), &comerAccount); err != nil {
			return
		}
		// oauthAccount不存在，创建并直接关联！
		if comerAccount.ID == 0 {
			comerAccount = model.ComerAccount{
				ComerID:   logonComerId,
				OIN:       oauth.GetUserID(),
				IsPrimary: true,
				Nick:      oauth.GetUserNick(),
				Avatar:    oauth.GetUserAvatar(),
				Type:      oauthType,
				IsLinked:  true,
			}
			if err = model.CreateAccount(mysql.DB, &comerAccount); err != nil {
				return
			}
			loginResponse.OauthLinked = true
			return nil, loginResponse
		} else if /*未关联Comer*/ comerAccount.ComerID == 0 {
			//
			if err = model.BindComerAccountToComerId(mysql.DB, comerAccount.ID, logonComerId); err != nil {
				return
			}
			loginResponse.OauthLinked = true
			return nil, loginResponse
		} else if /*关联到其他Comer了*/ comerAccount.ComerID != logonComerId {
			var anotherComer model.Comer
			if err = model.GetComerByID(mysql.DB, comerAccount.ComerID, &anotherComer); err != nil {
				return
			}
			/*其他Comer未绑定钱包,则oauth帐号可以换绑至此Comer*/
			if anotherComer.ID == 0 || anotherComer.Address == nil {
				if err = model.BindComerAccountToComerId(mysql.DB, comerAccount.ID, logonComerId); err != nil {
					return
				}
				return nil, loginResponse
			} else {
				err = errors.New(fmt.Sprintf("oauth has linked to anohter comer"))
				return
			}
		}
	}
	//////////
	return nil, loginResponse
}

func loginWithUnRegistredComer(oauth auth.OauthAccount, oauthType model.ComerAccountType) (err error, loginResponse account.OauthLoginResponse) {

	var (
		comerAccount model.ComerAccount
		comer        model.Comer
		comerProfile model.ComerProfile
		comerId      uint64
	)
	// 系统是否已存在该oauth的comerAccount
	if err = account.GetComerAccount(mysql.DB, model.GithubOauth, oauth.GetUserID(), &comerAccount); err != nil {
		return
	}
	loginResponse = account.OauthLoginResponse{}
	loginResponse.ComerID = 0
	loginResponse.IsProfiled = false
	loginResponse.OauthLinked = false
	loginResponse.OauthAccountId = comerAccount.ID
	// 首次登录
	if comerAccount.ID == 0 {
		if err = mysql.DB.Transaction(func(tx *gorm.DB) (erro error) {
			comer = model.Comer{}
			if erro = account.CreateComer(mysql.DB, &comer); erro != nil {
				return erro
			}
			comerAccount = model.ComerAccount{
				ComerID:   comer.ID,
				OIN:       oauth.GetUserID(),
				IsPrimary: true,
				Nick:      oauth.GetUserNick(),
				Avatar:    oauth.GetUserAvatar(),
				Type:      oauthType,
				IsLinked:  false,
			}
			if erro = account.CreateAccount(mysql.DB, &comerAccount); erro != nil {
				return erro
			}
			return nil
		}); err != nil {
			return
		}
	} else {
		// 2 已关联到comer
		if comerAccount.ComerID != 0 && comerAccount.IsLinked {
			if err = account.GetComerByID(mysql.DB, comerAccount.ComerID, &comer); err != nil {
				return
			}
			if comer.ID == 0 || comer.IsDeleted {
				err = errors.New(fmt.Sprintf("Comer does not exist or was deleted!"))
				return
			}
			if err = account.GetComerProfile(mysql.DB, comerAccount.ComerID, &comerProfile); err != nil {
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
	return
}
