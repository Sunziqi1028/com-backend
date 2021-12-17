package tags

//
//// GetStartupTagList return the all startup tags in list
//func GetStartupTagList() (response model.ListResponse, err error) {
//	tags, err := model.GetTagListByCategory(mysql.DB, model.StartupTag)
//	if err != nil {
//		return
//	}
//	list := make([]model.SingleTag, 0)
//	for _, tag := range tags {
//		t := model.SingleTag{
//			Name: tag.Name,
//			Code: tag.Code,
//		}
//		list = append(list, t)
//	}
//	response.Total = uint64(len(list))
//	response.List = list
//	return
//}
//
//// GetSkillTagList  return the all startup tags in list
//func GetSkillTagList() (response model.ListResponse, err error) {
//	tags, err := model.GetTagListByCategory(mysql.DB, model.SkillTag)
//	if err != nil {
//		return
//	}
//	for _, tag := range tags {
//		t := model.SingleTag{
//			Name: tag.Name,
//			Code: tag.Code,
//		}
//		response.List = append(response.List, t)
//	}
//	response.Total = uint64(len(response.List))
//	return
//}
