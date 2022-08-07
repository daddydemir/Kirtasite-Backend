package auth

import (
	"fmt"
	"github.com/daddydemir/kirtasiye-projesi/service"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(username string) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["iat"] = time.Now().Unix()
	claims["sub"] = username
	claims["exp"] = time.Now().Add(time.Hour * 10).Unix()

	tokenString, _ := token.SignedString([]byte("I-am-not-use-this-key"))

	return tokenString
}

func IsValid(data string) (bool, map[string]interface{}) {
	tkn, err := jwt.Parse(data, func(token *jwt.Token) (interface{}, error) {
		return []byte("I-am-not-use-this-key"), nil
	})
	if err != nil {
		return false, service.WrongTokenMessage()
	}

	if !tkn.Valid {
		return false, service.UnauthorizedMessage()
	} else {
		return true, nil
	}
}

func TokenParser(data string) string {
	claims := jwt.MapClaims{}
	_, _ = jwt.ParseWithClaims(data, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("I-am-not-use-this-key"), nil
	})
	username := fmt.Sprintf("%v", claims["sub"])
	return username
}
