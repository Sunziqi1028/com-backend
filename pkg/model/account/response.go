package account

import "ceres/pkg/model/tag"

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
	Skills []tag.Tag `json:"skills"`
}

// ComerOuterAccountListResponse response of the comer outer account list
type ComerOuterAccountListResponse struct {
	List  []ComerAccount `json:"list"`
	Total uint64         `json:"total"`
}
