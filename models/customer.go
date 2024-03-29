package models

type Customers struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	UserData Users  `json:"user" gorm:"foreignKey:user_id;references:id"`
}

type Customer struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
}
