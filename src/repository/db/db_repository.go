package db

import (
	"github.com/gocql/gocql"
	"github.com/implicithash/bookstore_oauth-api/src/clients/cassandra"
	"github.com/implicithash/bookstore_oauth-api/src/domain/access_token"
	"github.com/implicithash/bookstore_oauth-api/src/domain/access_token/utils/errors"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpires     = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(token access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	/*session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()*/

	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NotFoundError("no access token found with given id")
		}
		return nil, errors.InternalServerError(err.Error())
	}
	return &result, nil
}

func (r *dbRepository) Create(at access_token.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(queryCreateAccessToken,
		at.AccessToken,
		at.UserId,
		at.ClientId,
		at.Expires,
	).Exec(); err != nil {
		return errors.InternalServerError(err.Error())
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(queryUpdateExpires,
		at.Expires,
		at.AccessToken,
	).Exec(); err != nil {
		return errors.InternalServerError(err.Error())
	}
	return nil
}
