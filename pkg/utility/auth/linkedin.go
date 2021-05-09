package auth

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

/// LinkedIn REST Client
/// implemnetes the OauthClient interface
/// see https://docs.microsoft.com/en-us/linkedin/shared/authentication/authorization-code-flow?context=linkedin%2Fcontext&tabs=HTTPS
type LinkedIn struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	client       *http.Client
}

type linkedAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiredIn   uint64 `json:"expires_in"`
}

/// GetAccessToken
/// from the linkedIn open api the parameter in linkedin is named code
func (linkedin *LinkedIn) GetAccessToken(requestToken string) (accessToken string, err error) {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", requestToken)
	data.Set("client_id", linkedin.ClientID)
	data.Set("client_secret", linkedin.ClientSecret)
	data.Set("redirect_uri", linkedin.RedirectURI)
	request, _ := http.NewRequest(
		"POST",
		"https://www.linkedin.com/oauth/v2/accessToken",
		strings.NewReader(data.Encode()),
	)
	response, err := linkedin.client.Do(request)
	if err != nil {
		return
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	r := linkedAccessTokenResponse{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return
	}
	accessToken = r.AccessToken
	return
}

/// GetUserProfile
/// LinkedIn Oauth get user profile logic
/// see https://docs.microsoft.com/zh-cn/linkedin/shared/integrations/people/profile-api?context=linkedin/consumer/context
func (linkedin *LinkedIn) GetUserProfile(accessToken string) (account OauthAccount, err error) {

	return
}
