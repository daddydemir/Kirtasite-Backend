package models

type Price struct {
	Id           int     `json:"id"`
	StationeryId int     `json:"stationery_id" validate:"required"`
	Info         string  `json:"info" validate:"required,max=299"`
	Price        float32 `json:"price" validate:"required,gt=0"`
}
