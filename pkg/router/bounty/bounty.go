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
)

// CreateBounty create bounty
func CreateBounty(ctx *router.Context) {
	request := new(model.BountyRequest)
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
