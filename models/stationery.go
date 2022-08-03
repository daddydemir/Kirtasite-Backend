package models

type Stationery struct {
	UserId      int
	CompanyName string
	AddressId   int
	Score       float32
	UserData    User `gorm:"foreignKey:user_id;references:id"`
}
