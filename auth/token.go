package auth

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(username string) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["iat"] = time.Now().Unix()
	claims["sub"] = username
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	tokenString, _ := token.SignedString([]byte("I-am-not-use-this-key"))

	return tokenString
}

func IsValid(data string) (bool, map[string]string) {
	tkn, err := jwt.Parse(data, func(token *jwt.Token) (interface{}, error) {
		return []byte("I-am-not-use-this-key"), nil
	})
	if err != nil {
		return false, map[string]string{"message": "Token geçersiz"}
	}

	if !tkn.Valid {
		return false, map[string]string{"message": "Token'ın süresi dolmuş."}
	} else {
		return true, map[string]string{}
	}
}
