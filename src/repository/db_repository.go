package db_respository

import (
	"github.com/gocql/gocql"
	"github.com/sebagalan/bookstore_oauth-api/src/clients/cassandra"
	"github.com/sebagalan/bookstore_oauth-api/src/domain/access_token"
	"github.com/sebagalan/bookstore_users-api/utils/errors"
)

const (
	insertAccessToken = "INSERT INTO access_token (access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?)"
	getAccessToken    = "select access_token, user_id, client_id, expires from access_token where access_token = ? LIMIT 1"
	updateAccessToken = "update access_token set expires = ? where access_token = ?"
)

// DbRepository ...
type DbRepository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestError)
	Create(access_token.AccessToken) *errors.RestError
	Update(access_token.AccessToken) *errors.RestError
}

type dbRepository struct{}

func (dbr *dbRepository) GetByID(id string) (*access_token.AccessToken, *errors.RestError) {

	var result access_token.AccessToken
	session := cassandra.GetSession()

	err := session.Query(getAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserID,
		&result.ClientID,
		&result.Expires,
	)

	if err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("not access token found")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}

	return &result, nil
}

func (dbr *dbRepository) Create(accessToken access_token.AccessToken) *errors.RestError {

	session := cassandra.GetSession()

	err := session.Query(insertAccessToken,
		accessToken.AccessToken,
		accessToken.UserID,
		accessToken.ClientID,
		accessToken.Expires,
	).Exec()

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (dbr *dbRepository) Update(accessToken access_token.AccessToken) *errors.RestError {

	session := cassandra.GetSession()

	err := session.Query(updateAccessToken,
		accessToken.Expires,
		accessToken.AccessToken,
	).Exec()

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

// NewDbRepository ..
func NewDbRepository() DbRepository {
	return &dbRepository{}
}
