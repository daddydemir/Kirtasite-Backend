package models

type Address struct {
	Id         int    `json:"id"`
	CityId     int    `json:"city_id"`
	DistrictId int    `json:"district_id"`
	Name       string `json:"name"`
	CordinateX string `json:"cordinate_x"`
	CordinateY string `json:"cordinate_y"`
}
