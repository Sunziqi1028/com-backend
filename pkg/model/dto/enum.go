// @EgoctlOverwrite YES
// @EgoctlGenerateTime 20210223_200720
package dto

type EnumCreate struct {
	
	Id int `json:"id" binding:""` id
	
	GroupKey string `json:"groupKey" binding:""` 唯一标识符
	
	GroupTitle string `json:"groupTitle" binding:""` 名称
	
	Key int `json:"key" binding:""` 值
	
	Title string `json:"title" binding:""` 名称
	
	Ctime int64 `json:"ctime" binding:""` 创建时间
	
	Utime int64 `json:"utime" binding:""` 更新时间
	
	Dtime int64 `json:"dtime" binding:""` 删除时间
	
}

type EnumUpdate struct {
	
	Id int `json:"id" binding:""` id
	
	GroupKey string `json:"groupKey" binding:""` 唯一标识符
	
	GroupTitle string `json:"groupTitle" binding:""` 名称
	
	Key int `json:"key" binding:""` 值
	
	Title string `json:"title" binding:""` 名称
	
	Ctime int64 `json:"ctime" binding:""` 创建时间
	
	Utime int64 `json:"utime" binding:""` 更新时间
	
	Dtime int64 `json:"dtime" binding:""` 删除时间
	
}
