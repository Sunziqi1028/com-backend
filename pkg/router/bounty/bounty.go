/**
 * @Author: Sun
 * @Description:
 * @File:  bounty
 * @Version: 1.0.0
 * @Date: 2022/6/25 22:54
 */

package bounty

import (
	"ceres/pkg/model"
	"ceres/pkg/model/bounty"
	"ceres/pkg/router"
	"ceres/pkg/router/middleware"
	service "ceres/pkg/service/bounty"
	"fmt"
	"strconv"
)

// CreateBounty create bounty
func CreateBounty(ctx *router.Context) {
	request := new(bounty.BountyRequest)
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.HandleError(err)
		return
	}

	fmt.Println(request, "router/bounty.go line:26") // 注释

	if err := service.CreateComerBounty(request); err != nil {
		ctx.HandleError(err)
		return
	}
	response := "create bounty successful!"

	ctx.OK(response)
}

// GetPublicBountyList bounty list displayed in bounty tab
func GetPublicBountyList(ctx *router.Context) {
	var request model.Pagination
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.HandleError(err)
		return
	}
	request.Limit = 10

	if response, err := service.QueryAllBounties(request); err != nil {
		ctx.HandleError(err)
	} else {
		ctx.OK(response)
	}
}

// GetBountyListByStartup get bounty list belongs to startup
func GetBountyListByStartup(ctx *router.Context) {
	var request model.Pagination
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.HandleError(err)
		return
	}
	startupId, err := strconv.ParseUint(ctx.Param("accountID"), 0, 64)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	if startupId == 0 {
		err := router.ErrBadRequest.WithMsg("Invalid startupIdk")
		ctx.HandleError(err)
		return
	}
	request.Limit = 3

	if response, err := service.QueryBountiesByStartup(startupId, request); err != nil {
		ctx.HandleError(err)
	} else {
		ctx.OK(response)
	}
}

// GetMyPostedBountyList get bounty list posted by me
func GetMyPostedBountyList(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	var request model.Pagination
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.HandleError(err)
		return
	}
	request.Limit = 8

	if response, err := service.QueryComerPostedBountyList(comerID, request); err != nil {
		ctx.HandleError(err)
	} else {
		ctx.OK(response)
	}
}

// GetMyParticipatedBountyList get bounty list
func GetMyParticipatedBountyList(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	var request model.Pagination
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.HandleError(err)
		return
	}
	request.Limit = 8

	if response, err := service.QueryComerParticipatedBountyList(comerID, request); err != nil {
		ctx.HandleError(err)
	} else {
		ctx.OK(response)
	}
}
