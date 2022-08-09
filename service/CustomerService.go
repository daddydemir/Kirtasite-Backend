package service

import (
	"kirtasiteproject/auth"
	"kirtasiteproject/models"
	"kirtasiteproject/repositories"
	"strconv"
)

func AllCustomers(token string) (bool, map[string]interface{}) {
	status, message := auth.IsValid(token)
	if status {
		nm, durum := repositories.AllCustomers()
		if durum {
			message["data"] = nm
			return true, message
		} else {
			return false, message
		}
	} else {
		return false, message
	}
}

func GetCustomerByUserId(token string, userId string) (bool, map[string]interface{}) {
	status, message := auth.IsValid(token)
	if status {
		nm, durum := repositories.GetCustomerByUserId(userId)
		if durum {
			message["data"] = nm
			return true, message
		}
		return false, message
	} else {
		return false, message
	}
}

func UpdateCustomer(token string, user models.Customers) (bool, interface{}) {
	isValid, message := auth.IsValid(token)
	if isValid {
		isVal, newMessage := auth.UserAuth(token, user)
		if isVal {
			nm, status := repositories.UpdateCustomer(user)
			newMessage["data"] = nm
			if status {
				return true, newMessage
			}
			return false, nil
		} else {
			return false, newMessage
		}
	} else {
		return false, message
	}
}

func DeleteCustomer(token string, user models.Customers) (bool, interface{}) {
	isValid, message := auth.IsValid(token)
	if isValid {
		isVal, newMessage := auth.UserAuth(token, user)
		if isVal {
			nm, status := repositories.DeleteCustomer(strconv.Itoa(user.UserId))
			if status {
				newMessage["data"] = nm
				return true, newMessage
			}
			return false, nil
		} else {
			return false, newMessage
		}
	} else {
		return false, message
	}
}

func UpdateImageService(token string, user models.Customers) (bool, interface{}) {
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

func AddUserService(user models.Customers) (bool, interface{}) {
	m, s := repositories.AddCustomer(user)
	return s, m
}
