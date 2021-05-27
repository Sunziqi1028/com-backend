package tags

import (
	"ceres/pkg/initialization/mysql"
	model "ceres/pkg/model/tag"
)

/// GetStartupTagList
/// return the all startup tags in list
func GetStartupTagList() (response model.TagListResponse, err error) {
	tags, err := model.GetTagListByCategory(mysql.DB, model.StartupTag)
	if err != nil {
		return
	}
	var list []model.SingleTag
	for _, tag := range tags {
		t := model.SingleTag{
			Name: tag.Name,
			Code: tag.Code,
		}
		list = append(list, t)
	}
	response.Total = uint64(len(list))
	response.List = list
	return
}

/// GetSkillTagList
/// return the all startup tags in list
func GetSkillTagList() (response model.TagListResponse, err error) {
	tags, err := model.GetTagListByCategory(mysql.DB, model.SkillTag)
	if err != nil {
		return
	}
	var list []model.SingleTag
	for _, tag := range tags {
		t := model.SingleTag{
			Name: tag.Name,
			Code: tag.Code,
		}
		list = append(list, t)
	}
	response.Total = uint64(len(list))
	response.List = list
	return
}
