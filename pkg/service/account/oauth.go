package account

import "ceres/pkg/utility/auth"

/// common oauth login logic in comunion

func LoginWithOauth(client auth.OauthClient) (response interface{}, err error) {

	_, err = client.GetUserProfile()

	if err != nil {
		// Build the response object
		return
	}

	// do login logic

	// sign with jwt

	return
}
