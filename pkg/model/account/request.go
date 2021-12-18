package account

import (
	"ceres/pkg/initialization/utility"
	"ceres/pkg/router"
)

// EthLoginRequest the standard result of the web3.js signature
type EthLoginRequest struct {
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

func (v CreateProfileRequest) Validate() error {
	if len(v.Name) == 0 || len(v.Name) > 24 {
		return router.ErrBadRequest.WithMsg("Please enter your name")
	}
	if !utility.ValidateUrl(v.Website) {
		return router.ErrBadRequest.WithMsg("Please enter the correct network address")
	}
	if len(v.SKills) == 0 {
		return router.ErrBadRequest.WithMsg("Please enter your skill tag")
	}
	if len(v.BIO) < 100 {
		return router.ErrBadRequest.WithMsg("Please enter at least 100 characters")
	}
	return nil
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

func (v UpdateProfileRequest) Validate() error {
	if len(v.Name) == 0 || len(v.Name) > 24 {
		return router.ErrBadRequest.WithMsg("Please enter your name")
	}
	if !utility.ValidateUrl(v.Website) {
		return router.ErrBadRequest.WithMsg("Please enter the correct network address")
	}
	if len(v.SKills) == 0 {
		return router.ErrBadRequest.WithMsg("Please enter your skill tag")
	}
	if len(v.BIO) < 100 {
		return router.ErrBadRequest.WithMsg("Please enter at least 100 characters")
	}
	return nil
}
