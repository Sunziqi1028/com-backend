package auth

import (
	"ceres/pkg/config/auth"
	"crypto/tls"
	"net/http"
)

/// Comunion Oauth interface
/// Comunion Ceres only do the final legged in all Oauth2 processing
/// The Frontend will handle the other two legged using the standard Oauth2 API

/// OauthAccount
/// Oauth account interface to get the Oauth user unique ID nick name and the avatar
type OauthAccount interface {

	/// GetUserID
	/// get the user unique ID for every userID
	GetUserID() string

	/// GetUserNick
	/// get user nick name from Oauth Account
	GetUserNick() string

	/// GetUserAvatar
	/// get user avatar from Oauth Account
	GetUserAvatar() string
}

/// OauthClient
/// Abstraction of Oauth Login logic
type OauthClient interface {
	/// GetUserProfile
	GetUserProfile() (account OauthAccount, err error)
}

/// FIXME：should replace with ceres http library
var httpClient = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	},
}

/// NewGithubOauthClient
/// build a new Github client with the request token from login
func NewGithubOauthClient(requestToken string) (client OauthClient) {
	return &Github{
		ClientID:     auth.GithubClientID,
		ClientSecret: auth.GithubClientSecret,
		client:       httpClient,
		requestToken: requestToken,
	}
}

/// NewFacebookClient
/// build a new Facebook client with the request token from login 
func NewFacebookClient(requestToken string) (client OauthClient){
	return &Facebook{
		ClientID: auth.FacebookClientID,
		ClientSecret: auth.FacebookClientSecret,
		RedirectURI: auth.FacebookCallbackURL,
		client:  httpClient,
		RequestToken: requestToken,
	}
}

/// NewTwitterClient
/// build a new Twitter client with the request token from login 
func NewTwitterClient() (client OauthClient){

	return
}

/// NewLinkedinClient
/// build a new LinkedIn client with the request token from login 
func NewLinkedinClient() (client OauthClient){

	return
}