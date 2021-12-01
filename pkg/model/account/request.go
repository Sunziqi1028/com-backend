package account

// EthSignatureObject the standard result of the web3.js signature
// the signature use the spec256k1 algos
type EthSignatureObject struct {
	Address   string `json:"address"`
	Signature string `json:"signature"`
}

// CreateProfileRequest create a new profile then will let the entity to backend
type CreateProfileRequest struct {
	Name        string    		`json:"column:name"`
	Location    string    		`json:"column:location"`
	Website     string    		`json:"column:website"`
	Bio         string    		`json:"column:bio"`
	Socials     []SocialEntity  `json:"socials"`
	Skills      []string 		`json:"skills"`
	Wallets     []string 		`json:"wallets"`
}

// UpdateProfileRequest  update the comer profile
type UpdateProfileRequest struct {
	Name        string    		`json:"column:name"`
	Location    string    		`json:"column:location"`
	Website     string    		`json:"column:website"`
	Bio         string    		`json:"column:bio"`
	Socials     []SocialEntity  `json:"socials"`
	Skills      []string 		`json:"skills"`
	Wallets     []string 		`json:"wallets"`
}
