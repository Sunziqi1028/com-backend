package crowdfunding

import (
	crowdfunding2 "ceres/pkg/model/crowdfunding"
	"ceres/pkg/router"
	"ceres/pkg/router/middleware"
	"ceres/pkg/service/crowdfunding"
)

func CreateCrowdFunding(ctx *router.Context) {
	var request crowdfunding2.CreateCrowdfundingRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.HandleError(err)
		return
	}
	if err := request.ValidRequest(); err != nil {
		ctx.HandleError(err)
		return
	}
	request.ComerId = comerId(ctx)
	if err := crowdfunding.CreateCrowdfunding(request); err != nil {
		ctx.HandleError(err)
		return
	}

}

func SelectNonFundingStartups(ctx *router.Context) {
	comerId := comerId(ctx)
	startups, err := crowdfunding.SelectNonFundingStartups(comerId)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(startups)
}

func comerId(ctx *router.Context) uint64 {
	return ctx.Keys[middleware.ComerUinContextKey].(uint64)
}
