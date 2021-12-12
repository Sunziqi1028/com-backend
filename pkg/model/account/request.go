package account

// EthSignatureObject the standard result of the web3.js signature
// the signature use the spec256k1 algos
type EthSignatureObject struct {
	Address   string `json:"address"`
	Signature string `json:"signature"`
}

// CreateProfileRequest create a new profile then will let the entity to backend
type CreateProfileRequest struct {
	Name     string   `json:"name"`
	Avatar   string   `json:"avatar"`
	Location string   `json:"location"`
	Website  string   `json:"website"`
	SKills   []string `json:"skills"`
	BIO      string   `json:"bio"`
}

// UpdateProfileRequest  update the comer profile
type UpdateProfileRequest struct {
	Name     string   `json:"name"`
	Avatar   string   `json:"avatar"`
	Location string   `json:"location"`
	Website  string   `json:"website"`
	SKills   []string `json:"skills"`
	BIO      string   `json:"bio"`
}
