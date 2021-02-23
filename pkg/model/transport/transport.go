// @EgoctlOverwrite YES
// @EgoctlGenerateTime 20210223_202936
package transport

type ReqPage struct {
	Current  int    `json:"current" form:"current"`
	PageSize int    `json:"pageSize" form:"pageSize"`
	Sort     string `json:"sort" form:"sort"`
}