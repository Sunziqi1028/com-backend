package auth

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
/// Abstraction of the comunion oauth account login logic
type OauthClient interface {

	/// GetAccessToken
	GetAccessToken(requestToken string) (token string, err error)

	/// GetUserProfile
	GetUserProfile(accessToken string, userId string) (account OauthAccount, err error)
}
