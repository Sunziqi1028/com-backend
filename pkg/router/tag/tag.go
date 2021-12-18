package tag

import (
	"ceres/pkg/model/tag"
	"ceres/pkg/router"
	service "ceres/pkg/service/tag"
	"fmt"
)

// GetTagList get tag list
func GetTagList(ctx *router.Context) {
	var request tag.ListRequest
	if err := ctx.BindQuery(&request); err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid data format")
		ctx.HandleError(err)
		return
	}
	fmt.Println(request)
	var response tag.ListResponse
	if err := service.GetStartupTagList(request, &response); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(response)
}
