package account

// ComerLoginResponse comer login response
type ComerLoginResponse struct {
	ComerID    uint64 `json:"comerID"`
	Nick       string `json:"nick"`
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
}

// ComerOuterAccountListResponse response of the comer outer account list
type ComerOuterAccountListResponse struct {
	List  []ComerAccount `json:"list"`
	Total uint64         `json:"total"`
}

type GetComerInfoResponse struct {
	Comer
	ComerProfile  ComerProfile     `json:"comerProfile"`
	Follows       []FollowRelation `json:"follows"`
	FollowsCount  int64            `gorm:"-" json:"followsCount"`
	Followed      []FollowRelation `json:"followed"`
	FollowedCount int64            `gorm:"-" json:"followedCount"`
}

type IsFollowedResponse struct {
	IsFollowed bool `json:"isFollowed"`
}
