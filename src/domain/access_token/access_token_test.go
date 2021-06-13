package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

/*
go test  github.com/sebagalan/bookstore_oauth-api/src/domain/access_token
*/

func TestAccessTokenConstant(t *testing.T) {
	if expirationTime != 24 {
		t.Error("expirationTime constant should be 24")
	}
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()

	assert.False(t, at.isExpired(), "at shouldn't be expired")
	assert.True(t, at.UserID == 0, "new access token should'n a user asssociated")
	assert.EqualValues(t, "", at.AccessToken, "new access token should'n a user asssociated")

}

func TestAccessTokenIsExpired(t *testing.T) {

	at := AccessToken{}

	assert.True(t, at.isExpired(), "empty access token shouln't be expired")
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()

	assert.False(t, at.isExpired(), "access token expiring three our from now should not be expired")

}
