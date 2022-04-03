package service

import (
	"demir/auth"
	"demir/models"
)

func GetAllPricesService(token string) (bool, map[string]string) {
	status, message := auth.IsValid(token)
	if status {
		return true, message
	} else {
		return false, message
	}
}

func PriceByIdService(token string) (bool, map[string]string) {
	status, message := auth.IsValid(token)
	if status {
		return true, message
	} else {
		return false, message
	}
}

func PriceAddService(token string, price models.Price) (bool, map[string]string) {
	isValid, message := auth.IsValid(token)
	if isValid {
		isVal, newMessage := auth.PriceAuth(token, price)
		if isVal {
			return true, map[string]string{"message": "Ekleme başarılı."}
		} else {
			return false, newMessage
		}
	} else {
		return false, message
	}
}

func PriceUpdateService(token string, price models.Price) (bool, map[string]string) {
	isValid, message := auth.IsValid(token)
	if isValid {
		isVal, newMessage := auth.PriceAuth(token, price)
		if isVal {
			return true, map[string]string{"message": "Güncelleme başarılı."}
		} else {
			return false, newMessage
		}
	} else {
		return false, message
	}
}

func PriceDeleteService(token string, price models.Price) (bool, map[string]string) {
	isValid, message := auth.IsValid(token)
	if isValid {
		isVal, newMessage := auth.PriceAuth(token, price)
		if isVal {
			return true, map[string]string{"message": "Silme başarılı."}
		} else {
			return false, newMessage
		}
	} else {
		return false, message
	}
}
