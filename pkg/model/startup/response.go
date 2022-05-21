package startup

type ListStartupsResponse struct {
	List  []Startup `json:"list"`
	Total int64     `json:"total"`
}

type GetStartupResponse struct {
	Startup
}

type ExistStartupResponse struct {
	IsExist bool `json:"isExist"`
}

type FollowedByMeResponse struct {
	IsFollowed bool `json:"isFollowed"`
}
