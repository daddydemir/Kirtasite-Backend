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

type Orders struct {
	Id             int          `json:"id"`
	FileId         int          `json:"file_id"`
	CustomerId     int          `json:"customer_id"`
	StationeryId   int          `json:"stationery_id"`
	StatusId       int          `json:"status_id"`
	TotalPrice     float32      `json:"total_price"`
	CreatedDate    time.Time    `json:"created_date"`
	DeliveryDate   time.Time    `json:"delivery_date"`
	FileData       File         `json:"file" gorm:"foreignKey:file_id;references:id"`
	CustomerData   Customers    `json:"customer" gorm:"foreignKey:customer_id;references:user_id"`
	StationeryData Stationeries `json:"stationery" gorm:"foreignKey:stationery_id;references:id"`
	StatusData     Order_Status `json:"status" gorm:"foreignKey:status_id;references:id"`
}
