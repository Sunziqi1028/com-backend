package account

// ComerLoginResposne
type ComerLoginResponse struct {
	Nick    string `json:"nick"`
	Avatar  string `json:"avatar"`
	ComerID string `json:"comer_id"`
	Address string `json:"address"`
	Token   string `json:"token"`
	UIN     uint64 `json:"uin"`
}

// WalletNonceResponse wrap the nonce for formating rule in resposne
type WalletNonceResponse struct {
	Nonce string `json:"nonce"`
}

// ComerProfileResponse return the profile of some comer
type ComerProfileResponse struct {
	Name        string    `gorm:"column:name"`
	Location    string    `gorm:"column:location"`
	Website     string    `gorm:"column:website"`
	Bio         string    `gorm:"column:bio"`
	Socials     []SocialEntity `json:"socials"`
	Skills      []string `json:"skills"`
	Wallets     []string `json:"wallets"`
}

// ComerOuterAccountObject comer outer account object
type ComerOuterAccountObject struct {
	Identifier uint64 `json:"identifier"`
	UIN        uint64 `json:"uin"`
	OIN        string `json:"oin"`
	IsMain     bool   `json:"main"`
	Nick       string `json:"nick"`
	Avatar     string `json:"avatar"`
	Category   int    `json:"category"`
	Type       int    `json:"type"`
	IsLinked   bool   `json:"linked"`
}

// ComerOuterAccountListResponse response of the comer outer account list
type ComerOuterAccountListResponse struct {
	List  []ComerOuterAccountObject `json:"list"`
	Total uint64                    `json:"total"`
}
