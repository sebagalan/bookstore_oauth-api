package access_token

import (
	"strings"
	"time"

	"github.com/sebagalan/bookstore_users-api/utils/errors"
)

const (
	expirationTime = 24
)

//AccessToken ...
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() *errors.RestError {

	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if len(at.AccessToken) == 0 {
		return errors.NewBadRequestError("invalid access token")
	}

	if at.ClientID <= 0 {
		return errors.NewBadRequestError("invalid client id for access token")
	}

	if at.UserID <= 0 {
		return errors.NewBadRequestError("invalid user id for access token")
	}

	return nil

}

//GetNewAccessToken ...
func GetNewAccessToken() AccessToken {

	return AccessToken{
		UserID:  0,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) isExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
