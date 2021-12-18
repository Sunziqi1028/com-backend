package tag

import "ceres/pkg/router"

type ListRequest struct {
	IsIndex bool `form:"isIndex"`
	Limit   int  `form:"limit"`
	Offset  int  `form:"offset"`
}

func (l ListRequest) Validate() error {
	if l.Limit <= 0 || l.Limit >= 20 {
		return router.ErrBadRequest.WithMsg("please input right limit")
	}
	if l.Offset <= 0 {
		return router.ErrBadRequest.WithMsg("please input right offset")
	}
	return nil
}
