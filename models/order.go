package models

import "time"

type Order struct {
	Id               int       `json:"id"`
	FileId           int       `json:"file_id"`
	UserId           int       `json:"user_id"`
	StationeryId     int       `json:"stationery_id"`
	DeliveryDate     time.Time `json:"delivery_date"`
	VerificationCode string    `json:"verification_code"`
	TotalPrice       float32   `json:"total_price"`
}
