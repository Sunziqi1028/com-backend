package tags

import (
	"ceres/pkg/router"
	service "ceres/pkg/service/tags"
)

// GetStartupTagList  get the startup tag list
func GetStartupTagList(ctx *router.Context) {
	response, err := service.GetStartupTagList()
	if err != nil {
		ctx.ERROR(
			router.ErrBuisnessError,
			err.Error(),
		)
		return
	}
	ctx.OK(response)
}

// GetSkillTagList  get the skill tag list of the comer profile
func GetSkillTagList(ctx *router.Context) {
	response, err := service.GetSkillTagList()
	if err != nil {
		ctx.ERROR(
			router.ErrBuisnessError,
			err.Error(),
		)
		return
	}
	ctx.OK(response)
}
