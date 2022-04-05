package repositories

import (
	"github.com/daddydemir/kirtasiye-projesi/config"
	"github.com/daddydemir/kirtasiye-projesi/models"
	"github.com/daddydemir/kirtasiye-projesi/security"
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

func UpdateImage(url string, userId string) {
	var user models.User = GetUserById(userId)
	user.ImagePath = url
	config.DB.Save(&user)
}

func UserByName(username string) (models.User, bool) {
	var user models.User
	config.DB.Find(&user, "username = ?", username)
	if user.Username == "" {
		return models.User{}, false
	} else {
		return user, true
	}
}
