package access_token

import (
	"fmt"
	"github.com/implicithash/bookstore_oauth-api/src/domain/access_token/utils/errors"
	"github.com/implicithash/bookstore_users-api/utils/crypto_utils"
	"strings"
	"time"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// Used for password grant type
	Username string `json:"username"`
	Password string `json:"password"`
	// Used for client credentials grant type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (at *AccessTokenRequest) Validate() *errors.RestErr {
	switch at.GrantType {
	case grantTypePassword:
		break
	case grantTypeClientCredentials:
		break
	default:
		return errors.BadRequestError("invalid grant type parameter")
	}
	// TODO: Validate parameters for each grant_type
	return nil
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.BadRequestError("invalid access token id")
	}
	if at.UserId <= 0 {
		return errors.BadRequestError("invalid user id")
	}
	if at.ClientId <= 0 {
		return errors.BadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return errors.BadRequestError("invalid expiration time")
	}
	return nil
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId: userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	//now := time.Now().UTC()
	//expirationTime := time.Unix(at.Expires, 0)
	//return expirationTime.Before(now)
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}
