package startup

import (
	model "ceres/pkg/model/startup"
	"ceres/pkg/router"
	"ceres/pkg/router/middleware"
	service "ceres/pkg/service/startup"
	"strconv"

	"github.com/qiniu/x/log"
)

// ListStartups get startup list
func ListStartups(ctx *router.Context) {
	var request model.ListStartupRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		log.Warn(err)
		err = router.ErrBadRequest.WithMsg(err.Error())
		ctx.HandleError(err)
		return
	}

	var response model.ListStartupsResponse
	if err := service.ListStartups(0, &request, &response); err != nil {
		ctx.HandleError(err)
		return
	}
	for i, startup := range response.List {
		response.List[i].MemberCount = len(startup.Members)
		response.List[i].FollowCount = len(startup.Follows)
	}
	ctx.OK(response)
}

// ListStartupsMe get my startup list
func ListStartupsMe(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	var request model.ListStartupRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		log.Warn(err)
		err = router.ErrBadRequest.WithMsg(err.Error())
		ctx.HandleError(err)
		return
	}

	var response model.ListStartupsResponse
	if err := service.ListStartups(comerID, &request, &response); err != nil {
		ctx.HandleError(err)
		return
	}
	for i, startup := range response.List {
		response.List[i].MemberCount = len(startup.Members)
		response.List[i].FollowCount = len(startup.Follows)
	}

	ctx.OK(response)
}

// GetStartup get startup
func GetStartup(ctx *router.Context) {
	startupID, err := strconv.ParseUint(ctx.Param("startupID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid startup ID")
		ctx.HandleError(err)
		return
	}
	var response model.GetStartupResponse
	if err := service.GetStartup(startupID, &response); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(response)
}

// StartupNameIsExist get startup name is exist
func StartupNameIsExist(ctx *router.Context) {
	name := ctx.Param("name")
	if name == "" {
		err := router.ErrBadRequest.WithMsg("Startup's name has been used")
		ctx.HandleError(err)
		return
	}
	isExist, err := service.StartupNameIsExist(name)
	if err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(model.ExistStartupResponse{IsExist: isExist})
}

// StartupTokenContractIsExist get startup name is exist
func StartupTokenContractIsExist(ctx *router.Context) {
	name := ctx.Param("tokenContract")
	if name == "" {
		err := router.ErrBadRequest.WithMsg("Startup's token contract has been used")
		ctx.HandleError(err)
		return
	}

	isExist, err := service.StartupTokenContractIsExist(name)
	if err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(model.ExistStartupResponse{IsExist: isExist})
}

// FollowStartup follow Startup
func FollowStartup(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	startupID, err := strconv.ParseUint(ctx.Param("startupID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid startup ID")
		ctx.HandleError(err)
		return
	}
	if err = service.FollowStartup(comerID, startupID); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(nil)
}

// UnfollowStartup unfollow Startup
func UnfollowStartup(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	startupID, err := strconv.ParseUint(ctx.Param("startupID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid startup ID")
		ctx.HandleError(err)
		return
	}
	if err = service.UnfollowStartup(comerID, startupID); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(nil)
}

// ListFollowStartups list follow startup
func ListFollowStartups(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	var request model.ListStartupRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		log.Warn(err)
		err = router.ErrBadRequest.WithMsg(err.Error())
		ctx.HandleError(err)
		return
	}

	var response model.ListStartupsResponse
	if err := service.ListFollowStartups(comerID, &request, &response); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(response)
}

// ListParticipateStartups list participate startup
func ListParticipateStartups(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	var request model.ListStartupRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		log.Warn(err)
		err = router.ErrBadRequest.WithMsg(err.Error())
		ctx.HandleError(err)
		return
	}

	var response model.ListStartupsResponse
	if err := service.ListParticipateStartups(comerID, &request, &response); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(response)
}

// ListBeMemberStartups list as a member of startup
func ListBeMemberStartups(ctx *router.Context) {
	comerID, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid comer ID")
		ctx.HandleError(err)
		return
	}
	var request model.ListStartupRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		log.Warn(err)
		err = router.ErrBadRequest.WithMsg(err.Error())
		ctx.HandleError(err)
		return
	}
	var response model.ListStartupsResponse
	if err := service.ListBeMemberStartups(comerID, &request, &response); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(response)
}

// StartupFollowedByMe get startup is followed by me
func StartupFollowedByMe(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	startupID, err := strconv.ParseUint(ctx.Param("startupID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid startup ID")
		ctx.HandleError(err)
		return
	}
	isFollowed, err := service.StartupFollowedByComer(startupID, comerID)
	if err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(model.FollowedByMeResponse{IsFollowed: isFollowed})
}
