package db

import (
	"github.com/implicithash/bookstore_oauth-api/src/domain/access_token"
	"github.com/implicithash/bookstore_oauth-api/src/domain/access_token/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {

}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.InternalServerError("database connection not implemented yet")
}

