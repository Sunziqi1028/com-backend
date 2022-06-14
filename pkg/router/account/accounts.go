package account

import (
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/model/account"
	model "ceres/pkg/model/account"
	"ceres/pkg/router"
	"ceres/pkg/router/middleware"
	service "ceres/pkg/service/account"
	"ceres/pkg/utility/jwt"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

// UserInfo list all accounts of the Comer
func UserInfo(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	var response account.ComerLoginResponse
	if err := service.UserInfo(comerID, &response); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(response)
}

// ListAccounts list all accounts of the Comer
func ListAccounts(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	var response model.ComerOuterAccountListResponse
	if err := service.GetComerAccounts(comerID, &response); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(response)
}

// UnlinkAccount unlink accounts for the Comer
func UnlinkAccount(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	accountID, err := strconv.ParseUint(ctx.Param("accountID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid account ID")
		ctx.HandleError(err)
		return
	}
	err = service.UnlinkComerAccount(comerID, accountID)
	if err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(nil)
}

// LinkWithWallet link current account with wallet
func LinkWithWallet(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	var ethLoginRequest model.EthLoginRequest
	if err := ctx.BindJSON(&ethLoginRequest); err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid data format")
		ctx.HandleError(err)
		return
	}

	if err := service.LinkEthAccountToComer(comerID, ethLoginRequest.Address, ethLoginRequest.Signature); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(nil)
}

// GetComerInfo get comer
func GetComerInfo(ctx *router.Context) {
	comerID, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid comer ID")
		ctx.HandleError(err)
		return
	}
	var response model.GetComerInfoResponse
	if err := service.GetComerInfo(comerID, &response); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(response)
}

// GetComerInfoByAddress get comer by address
func GetComerInfoByAddress(ctx *router.Context) {
	address := ctx.Param("address")
	if address == "" {
		err := router.ErrBadRequest.WithMsg("Comer's address required")
		ctx.HandleError(err)
		return
	}
	var response model.GetComerInfoResponse
	if err := service.GetComerInfoByAddress(address, &response); err != nil {
		ctx.HandleError(err)
		return
	}

	if response.Comer.ID == 0 {
		ctx.OK(nil)
	} else {
		ctx.OK(response)
	}
}

// FollowComer follow Comer
func FollowComer(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	targetComerID, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid comer ID")
		ctx.HandleError(err)
		return
	}
	if comerID == targetComerID {
		err = router.ErrBadRequest.WithMsg("Can not follow myself")
		ctx.HandleError(err)
		return
	}
	if err = service.FollowComer(comerID, targetComerID); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(nil)
}

// UnfollowComer unfollow Comer
func UnfollowComer(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	targetComerID, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid comer ID")
		ctx.HandleError(err)
		return
	}
	if err = service.UnfollowComer(comerID, targetComerID); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(nil)
}

// ComerFollowedByMe get comer is followed by me
func ComerFollowedByMe(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	targetComerID, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid comer ID")
		ctx.HandleError(err)
		return
	}
	isFollowed, err := service.FollowedByComer(comerID, targetComerID)
	if err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(model.IsFollowedResponse{IsFollowed: isFollowed})
}

// OauthFirstLoginLinkedByWalletAddress oauth首次登录时候关联于钱包地址——钱包地址对应的comer不一定存在。不存在则创建！！
func OauthFirstLoginLinkedByWalletAddress(ctx *router.Context) {
	address := ctx.Query("address")
	if address == "" {
		handleError(ctx, errors.New("address required"))
		return
	}
	oauthAccountId, err := strconv.ParseUint(ctx.Query("oauthAccountId"), 10, 64)
	if oauthAccountId == 0 || err != nil {
		handleError(ctx, errors.New("oauth account id required"))
		return
	}
	var (
		comer           model.Comer
		currentAccount  model.ComerAccount
		existedAccounts []model.ComerAccount
		loginResponse   model.OauthLoginResponse
		profile         model.ComerProfile
	)
	if err := model.GetComerByAddress(mysql.DB, address, &comer); err != nil {
		handleError(ctx, err)
		return
	}

	if err := model.GetComerAccountById(mysql.DB, oauthAccountId, &currentAccount); err != nil {
		handleError(ctx, err)
		return
	}
	if currentAccount.ID == 0 {
		handleError(ctx, errors.New(fmt.Sprintf("oauth account with id %d does not exist", oauthAccountId)))
		return
	}

	// comer existed!
	if comer.ID != 0 {
		if currentAccount.ComerID != 0 && currentAccount.ComerID != comer.ID {
			handleError(ctx, errors.New("oauth account has bound to another comer"))
			return
		}

		if err := model.GetComerAccountsByComerId(mysql.DB, comer.ID, &existedAccounts); err != nil {
			handleError(ctx, err)
			return
		}

		if existedAccounts == nil || len(existedAccounts) == 0 {
			if err := service.LinkOauthToComer(oauthAccountId, comer.ID); err != nil {
				handleError(ctx, err)
				return
			}
		} else {
			var existed = false
			var sameWithCrtAccount = false
			for _, existedAccount := range existedAccounts {
				if existedAccount.ComerID != 0 && existedAccount.Type == currentAccount.Type {
					existed = true
					if existedAccount.ID == oauthAccountId {
						sameWithCrtAccount = true
					}
					break
				}
			}
			if existed && !sameWithCrtAccount {
				handleError(ctx, errors.New("the wallet has bound to another oauth account"))
				return
			}
			if !existed {
				if err := service.LinkOauthToComer(oauthAccountId, comer.ID); err != nil {
					handleError(ctx, err)
					return
				}
			}
		}
	} else /*comer does not exist*/ {
		comer = model.Comer{
			Address: &address,
		}
		if err := mysql.DB.Transaction(func(tx *gorm.DB) (er error) {
			if err := model.CreateComer(mysql.DB, &comer); err != nil {
				return err
			}
			if err := service.LinkOauthToComer(oauthAccountId, comer.ID); err != nil {
				return err
			}
			return
		}); err != nil {
			handleError(ctx, err)
			return
		}
	}
	var (
		isProfiled bool
		userName   = currentAccount.Nick
		avatar     = currentAccount.Avatar
	)
	if err := model.GetComerProfile(mysql.DB, comer.ID, &profile); err != nil {
		isProfiled = false
	} else if profile.ID == 0 {
		isProfiled = false
	} else {
		isProfiled = true
		userName = profile.Name
		avatar = profile.Avatar
	}
	loginResponse = model.OauthLoginResponse{
		ComerID:        comer.ID,
		Nick:           userName,
		Avatar:         avatar,
		Address:        *comer.Address,
		Token:          jwt.Sign(comer.ID),
		IsProfiled:     isProfiled,
		OauthLinked:    true,
		OauthAccountId: oauthAccountId,
	}
	ctx.OK(loginResponse)
	return
}
