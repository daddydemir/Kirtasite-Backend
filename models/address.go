package models

type Address struct {
	Id         int    `json:"id"`
	CityId     int    `json:"city_id"`
	DistrictId int    `json:"district_id"`
	Header     string `json:"header"`
	X          string `json:"x"`
	Y          string `json:"y"`
}
