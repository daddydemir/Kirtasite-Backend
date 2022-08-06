package models

import "time"

type Comment struct {
	Id           int       `json:"id"`
	CustomerId   int       `json:"customer_id"`
	StationeryId int       `json:"stationery_id"`
	Content      string    `json:"content"`
	CreatedDate  time.Time `json:"created_date"`
	Score        int       `json:"score"`
}

type Comments struct {
	Id             int          `json:"id"`
	CustomerId     int          `json:"customer_id"`
	StationeryId   int          `json:"stationery_id"`
	Content        string       `json:"content"`
	CreatedDate    time.Time    `json:"created_date"`
	Score          int          `json:"score"`
	CustomerData   Customers    `json:"customer" gorm:"foreignKey:customer_id;references:user_id"`
	StationeryData Stationeries `json:"stationery" gorm:"foreignKey:stationery_id;references:user_id"`
}
