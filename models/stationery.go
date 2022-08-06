package models

type Stationeries struct {
	UserId        int       `json:"user_id"`
	CompanyName   string    `json:"company_name"`
	AddressId     int       `json:"address_id"`
	Score         float32   `json:"score"`
	UserData      Users     `json:"user" gorm:"foreignKey:user_id;references:id"`
	AddressesData Addresses `json:"address" gorm:"foreignKey:address_id;references:id"`
}

type Stationery struct {
	UserId      int     `json:"user_id"`
	CompanyName string  `json:"company_name"`
	AddressId   int     `json:"address_id"`
	Score       float32 `json:"score"`
}
