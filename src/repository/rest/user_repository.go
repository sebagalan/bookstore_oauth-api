package rest

import (
	"github.com/sebagalan/bookstore_oauth-api/src/clients/cassandra"
	"github.com/sebagalan/bookstore_oauth-api/src/domain/access_token"
	"github.com/sebagalan/bookstore_users-api/utils/errors"
)


// RestRepository ...
type UserRepository interface {
	Loggin(string, string) (*users.User *errors.RestError)
}

type userRepository struct{}

// NewDbRepository ..
func NewDbRepository() UserRepository {
	return &userRepository{}
}
