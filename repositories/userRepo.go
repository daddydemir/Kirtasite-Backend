package repositories

import (
	"demir/config"
	"demir/models"
	//"gorm.io/gorm/clause"
)

func UserGetAll() []models.User {
	var users []models.User
	// ilişkilerini de getiriyor
	//config.DB.Preload(clause.Associations).Find(&users)
	config.DB.Find(&users)
	return users
}

func AddUser(user models.User) {
	config.DB.Create(&user)
}
