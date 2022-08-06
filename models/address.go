package models

type Address struct {
	Id         int    `json:"id"`
	CityId     int    `json:"city_id"`
	DistrictId int    `json:"district_id"`
	Header     string `json:"header"`
	X          string `json:"x"`
	Y          string `json:"y"`
}

type Addresses struct {
	Id           int      `json:"id"`
	CityId       int      `json:"city_id"`
	DistrictId   int      `json:"district_id"`
	Header       string   `json:"header"`
	X            string   `json:"x"`
	Y            string   `json:"y"`
	CityData     City     `json:"city" gorm:"foreignKey:city_id;references:id"`
	DistrictData District `json:"district" gorm:"foreignKey:district_id;references:id"`
}
