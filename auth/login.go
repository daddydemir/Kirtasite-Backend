package auth

import (
	"demir/repositories"
	"demir/security"
)

func LoginToSystem(username string, password string) (map[string]string, bool) {
	user, status := repositories.UserByName(username)
	if !status {
		return map[string]string{"message": "Kullanıcı adı yada Parola hatalı"}, false
	} else {
		isSuccess := security.CheckPassword(user.Password, password)
		if isSuccess {
			token := GenerateToken(user.Username)
			return map[string]string{"token": token}, true
		} else {
			return map[string]string{"message": "Kullanıcı adı yada Parola hatalı"}, false
		}
	}
}
