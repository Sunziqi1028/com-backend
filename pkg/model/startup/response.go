package startup

type ListStartupsResponse struct {
	List  []Startup `json:"list"`
	Total int64     `json:"total"`
}

type GetStartupResponse struct {
	Startup
}
