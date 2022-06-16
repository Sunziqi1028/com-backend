package account

type OauthLoginResponse struct {
	ComerID    uint64 `json:"comerID"`
	Nick       string `json:"nick"`
	Avatar     string `json:"avatar"`
	Address    string `json:"address"`
	Token      string `json:"token"`
	IsProfiled bool   `json:"isProfiled"`
	// OauthType wallet, oauthGithub, oauthGoogle
	// OauthType ComerAccountType `json:"oauthType"`
	// OauthLinked
	OauthLinked    bool   `json:"oauthLinked"`
	OauthAccountId uint64 `json:"oauthAccountId"`
}

// ComerLoginResponse comer login response
type ComerLoginResponse struct {
	ComerID       uint64                     `json:"comerID"`
	Nick          string                     `json:"nick"`
	Avatar        string                     `json:"avatar"`
	Address       string                     `json:"address"`
	Token         string                     `json:"token"`
	IsProfiled    bool                       `json:"isProfiled"`
	FirstLogin    bool                       `json:"firstLogin"`
	ComerAccounts []*OauthAccountBindingInfo `json:"comerAccounts"`
}

// WalletNonceResponse wrap the nonce for formating rule in resposne
type WalletNonceResponse struct {
	Nonce string `json:"nonce"`
}

// ComerProfileResponse return the profile of some comer
type ComerProfileResponse struct {
	*ComerProfile
	ComerAccounts []*OauthAccountBindingInfo `json:"comerAccounts"`
}
type OauthAccountBindingInfo struct {
	Linked      bool             `json:"linked"`
	AccountType ComerAccountType `json:"accountType"`
	AccountId   uint64           `json:"accountId"`
}

// ComerOuterAccountListResponse response of the comer outer account list
type ComerOuterAccountListResponse struct {
	List  []ComerAccount `json:"list"`
	Total uint64         `json:"total"`
}

type GetComerInfoResponse struct {
	Comer
	ComerProfile  ComerProfile    `json:"comerProfile"`
	Follows       []FollowComer   `json:"follows"`
	FollowsCount  int64           `json:"followsCount"`
	Followed      []FollowedComer `json:"followed"`
	FollowedCount int64           `json:"followedCount"`
}

type IsFollowedResponse struct {
	IsFollowed bool `json:"isFollowed"`
}

// LinkWalletResponse link wallet response
type LinkWalletResponse struct {
	IsProfiled bool `json:"isProfiled"`
}
