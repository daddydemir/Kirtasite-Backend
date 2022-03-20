package repositories

import (
	"demir/config"
	"demir/models"
	"demir/security"
	//"gorm.io/gorm/clause"
)

func GetUserById(userId string) models.User {
	var user models.User
	config.DB.Find(&user, "id = ?", userId)
	return user
}

func UserGetAll() []models.User {
	var users []models.User
	// ilişkilerini de getiriyor
	//config.DB.Preload(clause.Associations).Find(&users)
	config.DB.Find(&users)
	return users
}

func AddUser(user models.User) {
	user.Password = security.HashPassword(user.Password)
	config.DB.Create(&user)
}

func DeleteUser(user models.User) {
	config.DB.Delete(&user)
}

func UpdateUser(user models.User) {
	config.DB.Save(&user)
}
