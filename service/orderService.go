package service

import (
	"github.com/daddydemir/kirtasiye-projesi/auth"
	"github.com/daddydemir/kirtasiye-projesi/models"
)

func OrderByUserIdService(token string, userId int) (bool, map[string]string) {
	status, message := auth.IsValid(token)
	if status {
		isVal, newMessage := auth.OrderAuthUser(token, userId)
		if isVal {
			return true, map[string]string{"message": "İstek başarıyla işlendi."}
		} else {
			return false, newMessage
		}
	} else {
		return false, message
	}
}

func OrderByStationerIdService(token string, stationId int) (bool, map[string]string) {
	status, message := auth.IsValid(token)
	if status {
		isVal, newMessage := auth.OrderAuthStationery(token, stationId)
		if isVal {
			return true, map[string]string{"message": "İstek başarıyla işlendi."}
		} else {
			return false, newMessage
		}
	} else {
		return false, message
	}
}

func OrderAddService(token string, order models.Order) (bool, map[string]string) {
	status, message := auth.IsValid(token)
	if status {
		isVal, newMessage := auth.OrderAuth(token, order)
		if isVal {
			return true, map[string]string{"message": "Sipariş ekleme başarılı."}
		} else {
			return false, newMessage
		}
	} else {
		return false, message
	}
}

func OrderByIdServiceForUser(token string, order models.Order) (bool, map[string]string) {
	status, message := auth.IsValid(token)
	if status {
		isVal, newMessage := auth.OrderAuthUser(token, order.UserId)
		if isVal {
			return true, map[string]string{"message": "Sipariş ekleme başarılı."}
		} else {
			return false, newMessage
		}
	} else {
		return false, message
	}
}

func OrderByIdServiceForStationery(token string, order models.Order) (bool, map[string]string) {
	status, message := auth.IsValid(token)
	if status {
		isVal, newMessage := auth.OrderAuthStationery(token, order.StationeryId)
		if isVal {
			return true, map[string]string{"message": "Sipariş ekleme başarılı."}
		} else {
			return false, newMessage
		}
	} else {
		return false, message
	}
}
