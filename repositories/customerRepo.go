package repositories

import (
	"github.com/daddydemir/kirtasiye-projesi/config"
	"github.com/daddydemir/kirtasiye-projesi/models"
	"github.com/daddydemir/kirtasiye-projesi/security"
)

func returnModel(result error, model interface{}) (interface{}, bool) {
	if result == nil {
		return model, true
	} else {
		return models.User{}, false
	}
}

func GetCustomerByUserId(userId string) (interface{}, bool) {
	var customer models.Customers
	result := config.DB.Find(&customer, "user_id = ?", userId)
	return returnModel(result.Error, customer)
}

func GetCustomerByUserName(name string) (interface{}, bool) {
	var customer models.Customers
	result := config.DB.Find(&customer, "username = ?", name)
	return returnModel(result.Error, customer)
}

func AddCustomer(customer models.Customers) (interface{}, bool) {
	_, status := GetCustomerByUserName(customer.Username)
	if status {
		return "Bu isim zaten kullanılıyor.", false
	} else {
		customer.UserData.Password = security.HashPassword(customer.UserData.Password)
		result := config.DB.Create(&customer)
		return returnModel(result.Error, customer)
	}
}

func DeleteCustomer(userId string) (interface{}, bool) {
	customer, _ := GetCustomerByUserId(userId)
	result := config.DB.Delete(&customer)
	return returnModel(result.Error, "")
}

func UpdateCustomer(customer models.Customers) (interface{}, bool) {
	// TODO parola sıfıralama isteğini kontrol et
	result := config.DB.Save(&customer)
	return returnModel(result.Error, customer)
}

func AllCustomers() (interface{}, bool) {
	var customers []models.Customers
	result := config.DB.Find(&customers)
	return returnModel(result.Error, customers)
}

func UpdateCustomerImage(url string, userId string) (interface{}, bool) {
	response, _ := GetCustomerByUserId(userId)
	var customer models.Customers
	customer = response.(models.Customers)
	customer.UserData.ImagePath = url
	result := config.DB.Save(&customer)
	return returnModel(result.Error, customer)
}
