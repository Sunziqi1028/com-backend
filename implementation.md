# Implementation Notes 

Ceres every modules implementations notes will be placed in this documents.

## Framework(why ego)

The web development in go takes a big problem because of the simple syntax of go. Also in go there is not enough governance tools for projects, such as AOP and IOC. So we need a best practice framework to develop the Ceres. 

Ego, which is a web development tool kit in our opnion not a framework and more like a solution set of web devlopment which could give us the base tools like HTTP framework(Gin), ORM (gorm), RPC(gRpc), monitor, common logger...

## Project Structure

All the core code not only the API but also the DAO is placed in the pkg directory.

### Initialization behaviors 

Every framework work in Ceres will be written in the pkg/initialization package

### Model 

The models for database, currently we use Mysql as Ceres backend storeage. And every module define three files:

1. db.go which place the database models and dao logic
2. request.go request Data Transport Object for frontend JSON 
3. response.go response Object for HTTP APIs return 

### Rounter 

In pkg/router all the business services implemented in the inner directories. For exmple the account module implementes the Comer account logic 

* pkg/router/account/*.go is the implementations of business logic
* pkg/router/account.go is placed the all router about the account API registation function not only HTTP but gRpc

## Project Environment variables

1. CERES_MACHINE_IP to identify the machine IP to let the snowflake work correctly default is 127.0.0.1


## Comer 

Comer is more like the account in other websit such as Twitter, LinkedIn, Facebook, in concepts. But comunion does not save any critical informations of users such as password, login name. So we implement the comer with both Oauth and Web3 plugin to register or login. And we design the Web3 wallet address or Oauth user ID as unique identifier--UID but in Comunion all the user will use the identifier named UIN to identify a unique user. 

When user login in Comunion with Oauth or Web3, there are two branches to login:

1. if UID is not map to a Comer then create a new Comer with this UID account and then sign token to frontend
2. else use the existed Comer to sign the login token to frontend 

To abstract the Oauth behavior of many different websit we define two interface(which you could find in pkg/utility/auth/ auth.go) as below 

```go
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
	GetUserProfile(accessToken string) (account OauthAccount, err error)
}
```

In the account logic module we use two interface's implementations to login the account with two ways. The logic of login placed in pkg/router/account


## MISC

* middleware to catch the header Authorization carring the JWT token of comer if not is the Guest role 
* 

TODO: should complete soon 