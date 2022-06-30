/**
 * @Author: Sun
 * @Description:
 * @File:  bounty
 * @Version: 1.0.0
 * @Date: 2022/6/25 22:54
 */

package bounty

import (
	model "ceres/pkg/model/bounty"
	"ceres/pkg/router"
	service "ceres/pkg/service/bounty"
	"strconv"
)

// GetComerStartups get all startups of one comerID
func GetComerStartups(ctx *router.Context) {
	comerID, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid comer ID")
		ctx.HandleError(err)
		return
	}
	response, err := service.GetStartupsByComerID(comerID)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(response)
}

// CreateBounty create bounty
func CreateBounty(ctx *router.Context) {
	request := new(model.BountyRequest)
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.HandleError(err)
		return
	}
	if err := service.CreateComerBounty(request); err != nil {
		ctx.HandleError(err)
		return
	}
	response := &model.CreateBountyResponse{
		Data:   "create bounty successful!",
		Status: 0,
	}
	ctx.OK(response)
}

// GetPublicBountyList bounty list displayed in bounty tab
func GetPublicBountyList(ctx *router.Context) {

}

// GetBountyListByStartup get bounty list belongs to startup
func GetBountyListByStartup(ctx *router.Context) {

}

// GetMyPostedBountyList get bounty list posted by me
func GetMyPostedBountyList(ctx *router.Context) {

}

// GetMyParticipatedBountyList get bounty list
func GetMyParticipatedBountyList(ctx *router.Context) {

}

