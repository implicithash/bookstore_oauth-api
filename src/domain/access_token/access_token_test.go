package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T) {
	/*if expirationTime != 24 {
		t.Error("Expiration time should be 24 minutes")
	}*/
	assert.EqualValues(t, 24, expirationTime, "Expiration time should be 24 minutes")
}

func TestGetAccessToken(t *testing.T) {
	at := GetAccessToken()
	assert.False(t, at.IsExpired(), "brand new access token should be expired")
	assert.EqualValues(t, "", at.AccessToken, "new access token should not have defined access token id")
	assert.True(t, at.UserId == 0, "new access token should not have an associated user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(5 * time.Minute).Unix()
	assert.False(t, at.IsExpired(), "access token created 5 minutes from now should NOT be expired")
}
