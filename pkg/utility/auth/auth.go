package auth

import (
	"crypto/tls"
	"net/http"
)

/// the constraints will be injected at compile time with Github CI
/// see more at build.sh
/// FIXME：should move to github action ci
var (
	GithubOauthClientID     string
	GithubOauthClientSecret string
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
/// build a new Github Oauth Client with the request token from login
func NewGithubOauthClient(requestToken string) (client OauthClient) {
	return &Github{
		ClientID:     GithubOauthClientID,
		ClientSecret: GithubOauthClientSecret,
		client:       httpClient,
		requestToken: requestToken,
	}
}
