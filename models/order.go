package models

import "time"

type Order struct {
	Id               int       `json:"id"`
	FileId           int       `json:"file_id" validate:"required"`
	UserId           int       `json:"user_id" validate:"required"`
	StationeryId     int       `json:"stationery_id" validate:"required"`
	DeliveryDate     time.Time `json:"delivery_date"`
	VerificationCode string    `json:"verification_code"`
	TotalPrice       float32   `json:"total_price"`
}
