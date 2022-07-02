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
	"fmt"
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
	fmt.Println(comerID, "router/bounty.go line:27") // 注释
	responseTmp, err := service.GetStartupsByComerID(comerID)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	if len(responseTmp) == 0 {
		response := &model.GetStartupsResponse{
			Data:        "please to create startup!",
			GetStartups: nil,
		}
		ctx.OK(response)
		return
	}
	response := &model.GetStartupsResponse{
		Data:        "get startup success!",
		GetStartups: responseTmp,
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

	fmt.Println(request, "router/bounty.go line:43") // 注释

	if err := service.CreateComerBounty(request); err != nil {
		ctx.HandleError(err)
		return
	}
	response := "create bounty successful!"

	ctx.OK(response)
}
