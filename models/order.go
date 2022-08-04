package models

import "time"

type Order struct {
	Id           int       `json:"id"`
	FileId       int       `json:"file_id"`
	CustomerId   int       `json:"customer_id"`
	StationeryId int       `json:"stationery_id"`
	StatusId     int       `json:"status_id"`
	TotalPrice   float32   `json:"total_price"`
	CreatedDate  time.Time `json:"created_date"`
	DeliveryDate time.Time `json:"delivery_date"`
}
