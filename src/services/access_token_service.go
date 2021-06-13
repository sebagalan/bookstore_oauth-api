package services

import (
	"github.com/sebagalan/bookstore_oauth-api/src/domain/access_token"
	"github.com/sebagalan/bookstore_users-api/utils/errors"
)

// Repository ...
type Repository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestError)
	Create(access_token.AccessToken) *errors.RestError
	Update(access_token.AccessToken) *errors.RestError
}

// AccessTokenServices ...
type AccessTokenServices interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestError)
	Create(access_token.AccessToken) *errors.RestError
	Update(access_token.AccessToken) *errors.RestError
}

type accessTokenServices struct {
	repository Repository
}

// NewAccessTokenServices ...
func NewAccessTokenServices(repo Repository) AccessTokenServices {
	return &accessTokenServices{
		repository: repo,
	}
}

func (ats *accessTokenServices) GetByID(id string) (*access_token.AccessToken, *errors.RestError) {

	acccessToken, err := ats.repository.GetByID(id)

	if err != nil {
		return nil, err
	}
	return acccessToken, nil
}

func (ats *accessTokenServices) Create(at access_token.AccessToken) *errors.RestError {

	err := at.Validate()

	if err != nil {
		return err
	}

	ats.repository.Create(at)

	return nil
}

func (ats *accessTokenServices) Update(at access_token.AccessToken) *errors.RestError {

	err := at.Validate()

	if err != nil {
		return err
	}

	ats.repository.Update(at)

	return nil
}
