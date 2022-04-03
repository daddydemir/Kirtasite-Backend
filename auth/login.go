package auth

import (
	"demir/repositories"
	"demir/security"
)

func LoginToSystem(username string, password string) (map[string]string, bool) {
	user, status := repositories.UserByName(username)
	stationery, state := repositories.StationeryByName(username)
	if !status && !state {
		return map[string]string{"message": "Kullanıcı adı yada Parola hatalı.1"}, false
	} else if status {
		isSuccess := security.CheckPassword(user.Password, password)
		if isSuccess {
			token := GenerateToken(user.Username)
			return map[string]string{"token": token}, true
		} else {
			return map[string]string{"message": "Kullanıcı adı yada Parola hatalı.2"}, false
		}
	} else if state {
		isSuccess := security.CheckPassword(stationery.Password, password)
		if isSuccess {
			token := GenerateToken(stationery.CompanyName)
			return map[string]string{"token": token}, true
		} else {
			return map[string]string{"message": "Kullanıcı adı yada Parola hatalı.3"}, false
		}
	} else {
		return map[string]string{"message": "Beklenmedik bir hata oluştu, daha sonra tekrar deneyiniz"}, false
	}
}
