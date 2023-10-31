package jwt_auth

import (
	"bill/modules/setting"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	//Secret 用于jwt加密
	Secret           = []byte("fsrm.90.it.cg")
	Timeout          = time.Duration(time.Minute * 10)
	SigningAlgorithm = "HS256"
)

func init() {
	//Timeout = time.Duration(global.ConfigFile.GetInt("jwt.timeout")) * time.Minute
	Timeout = time.Duration(setting.GetConfig().Jwt.Timeout) * time.Minute
}

func NewToken(claims jwt.MapClaims) (string, error) {

	if _, ok := claims["exp"]; !ok {
		expire := time.Now().Add(Timeout).Unix()
		claims["exp"] = expire
	}

	if _, ok := claims["orig_iat"]; !ok {
		claims["orig_iat"] = time.Now().Unix()
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod(SigningAlgorithm), claims)

	tokenString, err := token.SignedString(Secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 验证token
func VerifyToken(tokenString string) (*jwt.Token, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return Secret, nil
	})
	if err != nil {
		return nil, false
	}

	return token, token.Valid
}

func VerifyTokenAsClaims(tokenString string) (jwt.MapClaims, bool) {
	token, ok := VerifyToken(tokenString)

	if !ok {
		return nil, false
	}

	claims, ok2 := token.Claims.(jwt.MapClaims)

	return claims, ok2

}
