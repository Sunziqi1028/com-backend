package account

/// EthSignatureObject
/// the standard result of the web3.js signature
/// the signature use the spec256k1 algos
type EthSignatureObject struct {
	Address     string `json:"address"`
	MessageHash string `json:"message_hash"`
	V           string `json:"v"`
	R           string `json:"r"`
	S           string `json:"s"`
	Signature   string `json:"signature"`
}

/// CreateProfileRequest
/// create a new profile then will let the entity to backend
type CreateProfileRequest struct {
	SKills      []uint64 `json:"skills"`
	Description string   `json:"description"`
	About       string   `json:"about"`
	Email       string   `json:"email"`
}

/// UpdateProfileRequest
/// update the comer profile
type UpdateProfileRequest struct {
	Identifier  uint64   `json:"identifier"`
	SKills      []uint64 `json:"skills"`
	Description string   `json:"description"`
	About       string   `json:"about"`
	Email       string   `json:"email"`
}
