package account

/// EthSignatureObject
/// the standard result of the web3.js signature
/// the signature use the spec256k1 algos
type EthSignatureObject struct {
	Address     string `json:"Address"`
	MessageHash string `json:"message_hash"`
	V           string `json:"v"`
	R           string `json:"r"`
	S           string `json:"s"`
	Signature   string `json:"signature"`
}
