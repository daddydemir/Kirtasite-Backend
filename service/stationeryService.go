package service

import (
	"github.com/daddydemir/kirtasiye-projesi/auth"
	"github.com/daddydemir/kirtasiye-projesi/models"
	"github.com/daddydemir/kirtasiye-projesi/repositories"
)

func GetAllStationeryService(token string) (bool, map[string]string) {
	status, message := auth.IsValid(token)
	if status {
		return true, message
	} else {
		return false, message
	}
}

func GetStationeryByIdService(token string) (bool, map[string]string) {
	status, message := auth.IsValid(token)
	if status {
		return true, message
	} else {
		return false, message
	}
}

func UpdateStationeryService(token string, stationery models.Stationery) (bool, map[string]string) {
	isValid, message := auth.IsValid(token)
	if isValid {
		isVal, newMessage := auth.StationeryAuth(token, stationery)
		if isVal {
			return true, map[string]string{"message": "Güncelleme başarılı."}
		} else {
			return false, newMessage
		}
	} else {
		return false, message
	}
}

func DeleteStationeryService(token string, stationery models.Stationery) (bool, map[string]string) {
	isValid, message := auth.IsValid(token)
	if isValid {
		isVal, newMessage := auth.StationeryAuth(token, stationery)
		if isVal {
			return true, map[string]string{"message": "Silme başarılı."}
		} else {
			return false, newMessage
		}
	} else {
		return false, message
	}
}

func UpdateSImageService(token string, stationery models.Stationery) (bool, map[string]string) {
	isValid, message := auth.IsValid(token)
	if isValid {
		isVal, newMessage := auth.StationeryAuth(token, stationery)
		if isVal {
			return true, map[string]string{"message": "Profil resmi başarıyla güncellendi."}
		} else {
			return false, newMessage
		}
	} else {
		return false, message
	}
}

func AddStationeryService(stationery models.Stationery) (bool, map[string]string) {
	_, status := repositories.StationeryByName(stationery.CompanyName)
	if status {
		return false, map[string]string{"message": "Bu isim zaten kullanılıyor."}
	} else {
		return true, map[string]string{"message": "Şirket ismi uygundur."}
	}
}
