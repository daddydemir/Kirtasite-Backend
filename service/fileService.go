package service

import (
	"demir/auth"
	"demir/models"
)

func GetFileByUserIdService(token string, userId int) (bool, map[string]string) {
	status, message := auth.IsValid(token)
	if status {
		isVal, newMessage := auth.FileAuthUser(token, userId)
		if isVal {
			return true, map[string]string{"message": "İstek başarıyla işlendi"}
		} else {
			return false, newMessage
		}
	} else {
		return false, message
	}
}

func FileByIdService(token string, fileId string) (bool, map[string]string) {
	status, message := auth.IsValid(token)
	if status {
		return true, map[string]string{"message": ""}
	} else {
		return false, message
	}
}

func FileDeleteService(token string, file models.File) (bool, map[string]string) {
	status, message := auth.IsValid(token)
	if status {
		isVal, newMessage := auth.FileAuthDelete(token, file)
		if isVal {
			return true, map[string]string{"message": "İstek başarıyla işlenedi."}
		} else {
			return false, newMessage
		}
	} else {
		return false, message
	}
}

func FileAddService(token string, file models.File) (bool, map[string]string) {
	status, message := auth.IsValid(token)
	if status {
		isVal, newMessage := auth.FileAuthUser(token, file.UserId)
		if isVal {
			return true, map[string]string{"message": "Dosya yüklendi."}
		} else {
			return false, newMessage
		}
	} else {
		return false, message
	}
}
