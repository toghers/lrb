package share

import (
	"fmt"

	"github.com/dgrijalva/jwt-go/v4"
)

type MyCustomClaims struct {
	Id int `json:"foo"`
	jwt.StandardClaims
}

func CreateTokens(claims *MyCustomClaims, key []byte) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	fmt.Printf("%v %v", ss, err)
	return ss
}
