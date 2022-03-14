package aggregate

import (
	"log"
	"testing"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/stretchr/testify/assert"
)

func Test_passportAggr_Login_jwt(t *testing.T) {
	claims := map[string]interface{}{
		"sub": "test",
	}
	jwtauth.SetIssuedNow(claims)
	jwtauth.SetExpiryIn(claims, time.Hour*3)

	_, tokenString, err := tokenAuth.Encode(claims)
	assert.Nil(t, err)

	log.Println(tokenString)
}
