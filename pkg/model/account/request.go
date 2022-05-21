package account

// EthLoginRequest the standard result of the web3.js signature
type EthLoginRequest struct {
	Address   string `json:"address" binding:"len=42,startswith=0x"`
	Signature string `json:"signature" binding:"required"`
}

// CreateProfileRequest create a new profile then will let the entity to backend
type CreateProfileRequest struct {
	Name     string   `json:"name" binding:"min=1,max=24"`
	Avatar   string   `json:"avatar" binding:"required"`
	Location string   `json:"location"`
	TimeZone *string  `json:"timeZone" binding:"required"`
	Website  string   `json:"website"`
	Email    *string  `json:"email" binding:"required"`
	SKills   []string `json:"skills" binding:"min=1"`
	Twitter  *string  `json:"twitter" binding:"required"`
	Discord  *string  `json:"discord" binding:"required"`
	Telegram *string  `json:"telegram" binding:"required"`
	Medium   *string  `json:"medium" binding:"required"`
	BIO      string   `json:"bio" binding:"min=100"`
}

// UpdateProfileRequest  update the comer profile
type UpdateProfileRequest struct {
	Name     string   `json:"name" binding:"min=1,max=24"`
	Avatar   string   `json:"avatar" binding:"required"`
	Location string   `json:"location"`
	TimeZone *string  `json:"timeZone" binding:"required"`
	Website  string   `json:"website"`
	Email    *string  `json:"email" binding:"required"`
	SKills   []string `json:"skills" binding:"min=1"`
	Twitter  *string  `json:"twitter" binding:"required"`
	Discord  *string  `json:"discord" binding:"required"`
	Telegram *string  `json:"telegram" binding:"required"`
	Medium   *string  `json:"medium" binding:"required"`
	BIO      string   `json:"bio" binding:"min=100"`
}
