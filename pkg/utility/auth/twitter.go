package auth

type Twitter struct {
	ConsumerKey    string
	ConsumerSecret string
}

/// TwitterOauthAccount
/// Twitter Oauth user profile account
type TwitterOauthAccount struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	ProfileImageURL string `json:"profile_image_url_https"`
}

/// GetUserID
func (account *TwitterOauthAccount) GetUserID() string {
	return account.ID
}

/// GetUserNick
func (account *TwitterOauthAccount) GetUserNick() string {
	return account.Name
}

/// GetUserAvatar
func (account *TwitterOauthAccount) GetUserAvatar() string {
	return account.ProfileImageURL
}
