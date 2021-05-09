package auth

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

/// Github REST client
/// implemnetes the OauthClient interface
type Github struct {
	ClientID     string
	ClientSecret string
	client       *http.Client
}

/// NewGithubClient
/// create a new github client
func NewGithubClient(clientID, clientSecret string) (github *Github) {
	github = &Github{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	}
	return
}

type githubAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

/// GetAccessToken
/// use the requestToken to get the access token which will be used to get the github user information
func (github *Github) GetAccessToken(requestToken string) (accessToken string, err error) {
	u, _ := url.Parse(
		fmt.Sprintf(
			"https://github.com/login/oauth/access_token?%s&%s&%s",
			github.ClientID,
			github.ClientSecret,
			requestToken,
		),
	)
	request := &http.Request{
		Method: "GET",
		URL:    u,
		Header: map[string][]string{
			"Accept": {"application/json"},
		},
	}
	resp, err := github.client.Do(request)
	if err != nil {
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	r := githubAccessTokenResponse{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return
	}
	accessToken = r.AccessToken
	return
}

/// GithubUserProfile
/// response of api.github.com/user
/// this struct implement the ceres/session/OauthAccount interface
type GithubUserProfile struct {
	Login  string `json:"login"`
	ID     uint64 `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar_url"`
}

/// implement the OauthAccount interface
func (account *GithubUserProfile) GetUserID() string {
	return account.Login
}

/// implement the OauthAccount interface
func (account *GithubUserProfile) GetUserAvatar() string {
	return account.Avatar
}

/// implement the OauthAccount interface
func (account *GithubUserProfile) GetUserNick() string {
	return account.Name
}

/// GetUserProfile
/// get user profile information from api.github.com/user
func (github *Github) GetUserProfile(accessToken string) (account OauthAccount, err error) {
	u, _ := url.Parse("https://api.github.com/user")
	request := &http.Request{
		Method: "GET",
		URL:    u,
		Header: map[string][]string{
			"Accept":        {"application/json"},
			"Authorization": {fmt.Sprintf("token %s", accessToken)},
		},
	}
	resp, err := github.client.Do(request)
	if err != nil {
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	json.Unmarshal(body, &account)
	return
}
