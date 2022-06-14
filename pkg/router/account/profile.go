package account

import (
	"ceres/pkg/initialization/mysql"
	model "ceres/pkg/model/account"
	"ceres/pkg/router"
	"ceres/pkg/router/middleware"
	service "ceres/pkg/service/account"
	"ceres/pkg/utility/auth"
	"fmt"
	"github.com/qiniu/x/log"
)

// CreateProfile create the profile
func CreateProfile(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	request := &model.CreateProfileRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		log.Warn(err)
		err = router.ErrBadRequest.WithMsg(err.Error())
		ctx.HandleError(err)
		return
	}

	if err := service.CreateComerProfile(comerID, request); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(nil)
}

// GetProfile get current Comer profile
func GetProfile(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	var response model.ComerProfileResponse
	if err := service.GetComerProfile(comerID, &response); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(response)
}

// UpdateProfile update the profile
func UpdateProfile(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	request := &model.UpdateProfileRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		log.Warn(err)
		err = router.ErrBadRequest.WithMsg(err.Error())
		ctx.HandleError(err)
		return
	}

	if err := service.UpdateComerProfile(comerID, request); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(nil)
}

func LinkOauth2Comer(ctx *router.Context) {
	comerID, err := extractComerIdFromJwtToken(ctx)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	var linkReq model.LinkOauth2WalletRequest
	if err := ctx.ShouldBindJSON(&linkReq); err != nil {
		ctx.HandleError(router.ErrBadRequest.WithMsg(err.Error()))
		return
	}
	var walletComer model.Comer
	if err := model.GetComerByID(mysql.DB, comerID, &walletComer); err != nil {
		ctx.HandleError(router.ErrBadRequest.WithMsg(fmt.Sprintf("Comer  %d does not exist", comerID)))
		return
	}
	if linkReq.WalletAddressRequired && walletComer.Address == nil {
		// 钱包对应的Comer必须存在！
		ctx.HandleError(router.ErrBadRequest.WithMsg(fmt.Sprintf("Comer  %d does not have wallet address", comerID)))
		return
	} else if walletComer.ID != comerID {
		// 当前登录人和传入的钱包必须是同一个Comer
		ctx.HandleError(router.ErrBadRequest.WithMsg(fmt.Sprintf("Invalid walletAddress: %d", walletComer.ID)))
		return
	}
	if linkReq.OauthType != model.GithubOauth && linkReq.OauthType != model.GoogleOauth {
		ctx.HandleError(router.ErrBadRequest.WithMsg(fmt.Sprintf("Invalid oauthType: %d", linkReq.OauthType)))
		return
	}
	//
	var targetOauthComer model.Comer
	var targetOauthComerAccount model.ComerAccount
	client := auth.NewGoogleClient(linkReq.OauthCode)
	if oauthAccount, err := client.GetUserProfile(); err != nil {
		ctx.HandleError(err)
		return
	} else {
		if err := model.GetComerAccount(mysql.DB, linkReq.OauthType, oauthAccount.GetUserID(), &targetOauthComerAccount); err != nil {
			walletComer = model.Comer{}
			return
		}
		model.GetComerByID(mysql.DB, targetOauthComerAccount.ComerID, &targetOauthComer)

	}

}
