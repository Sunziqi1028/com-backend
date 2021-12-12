package account

// ComerLoginResponse comer login response
type ComerLoginResponse struct {
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	Address    string `json:"address"`
	Token      string `json:"token"`
	IsProfiled bool   `json:"isProfiled"`
}

// WalletNonceResponse wrap the nonce for formating rule in resposne
type WalletNonceResponse struct {
	Nonce string `json:"nonce"`
}

// ComerProfileResponse return the profile of some comer
type ComerProfileResponse struct {
	ComerProfile
	Skills []Skill `json:"skills"`
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
