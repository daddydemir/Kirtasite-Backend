package service

import (
	"demir/auth"
	"demir/models"
	"demir/repositories"
)

func GetAllUserService(token string) (bool, map[string]string) {
	status, message := auth.IsValid(token)
	if status {
		return true, message
	} else {
		return false, message
	}
}

func GetUserByIdService(token string) (bool, map[string]string) {
	status, message := auth.IsValid(token)
	if status {
		return true, message
	} else {
		return false, message
	}
}

func UpdateUserService(token string, user models.User) (bool, map[string]string) {
	isValid, message := auth.IsValid(token)
	if isValid {
		isVal, newMessage := auth.UserAuth(token, user)
		if isVal {
			return true, map[string]string{"message": "Güncelleme başarılı."}
		} else {
			return false, newMessage
		}
	} else {
		return false, message
	}
}

func DeleteUserService(token string, user models.User) (bool, map[string]string) {
	isValid, message := auth.IsValid(token)
	if isValid {
		isVal, newMessage := auth.UserAuth(token, user)
		if isVal {
			return true, map[string]string{"message": "Silme başarılı."}
		} else {
			return false, newMessage
		}
	} else {
		return false, message
	}
}

func UpdateImageService(token string, user models.User) (bool, map[string]string) {
	isValid, message := auth.IsValid(token)
	if isValid {
		isVal, newMessage := auth.UserAuth(token, user)
		if isVal {
			return true, map[string]string{"message": "Profil resmi başarıyla güncellendi."}
		} else {
			return false, newMessage
		}
	} else {
		return false, message
	}
}

func AddUserService(user models.User) (bool, map[string]string) {
	_, status := repositories.UserByName(user.Username)
	if status {
		return false, map[string]string{"message": "Bu isme sahip bir kullanıcı zaten var."}
	} else {
		return true, map[string]string{"message": "Kullanıcı ismi uygundur."}
	}
}
