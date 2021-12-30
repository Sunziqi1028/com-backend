package startup

import (
	"ceres/pkg/model"
)

type ListStartupRequest struct {
	model.ListRequest
	Keyword string `form:"keyword"`
	Mode    Mode   `form:"mode"`
}
