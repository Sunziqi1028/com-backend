package crowdfunding

type CrowdfundableStartup struct {
	StartupId     uint64 `json:"startupId"`
	StartupName   string `json:"startupName"`
	CanRaise      bool   `json:"canRaise"`
	TokenContract string `json:"tokenContract,omitempty"`
}
