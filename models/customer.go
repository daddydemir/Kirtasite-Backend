package models

type Customer struct {
	UserId   int
	Username string
	UserData User `gorm:"foreignKey:user_id;references:id"`
}
