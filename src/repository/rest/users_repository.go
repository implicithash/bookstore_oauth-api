package rest

import (
	"encoding/json"
	"github.com/implicithash/bookstore_oauth-api/src/domain/access_token/utils/errors"
	"github.com/implicithash/bookstore_oauth-api/src/domain/users"
	"github.com/mercadolibre/golang-restclient/rest"
	_ "github.com/mercadolibre/golang-restclient/rest"
	"time"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "https://api.bookstore.com",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct {
}

func NewRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response := usersRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, errors.InternalServerError("invalid rest client response when trying to login user")
	}
	if response.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.InternalServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}
	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.InternalServerError("error when trying to unmarshal users login response")
	}
	return &user, nil
}
