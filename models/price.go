package models

type Price struct {
	Id           int     `json:"id"`
	StationeryId int     `json:"stationery_id"`
	Info         string  `json:"info"`
	Price        float32 `json:"price"`
}
