package jwt_auth

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
)

func TestNewToken(t *testing.T) {

	tokenStr, err := NewToken(jwt.MapClaims{
		"id":   1,
		"name": "kinson",
	})

	if err != nil {
		t.Error(err)
	}

	if _, ok := VerifyToken(tokenStr); !ok {
		t.Error("token无效")
	}
}
