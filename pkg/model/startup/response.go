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

type ListComerStartupsResponse struct {
	List  []*ListComerStartup `json:"list"`
	Total int                 `json:"total"`
}

type ListComerStartup struct {
	StartupID uint64 `gorm:"column:id" json:"startupID"`
	Name      string `gorm:"column:name" json:"name"`
}
