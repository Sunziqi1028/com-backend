package startup_team

import "ceres/pkg/model"

type ListStartupTeamMemberRequest struct {
	model.ListRequest
}

type CreateStartupTeamMemberRequest struct {
	ComerID  uint64 `json:"comerID" validate:"required"`
	Position string `json:"position" validate:"required"`
}

type UpdateStartupTeamMemberRequest struct {
	Position string `json:"position" validate:"required"`
}
