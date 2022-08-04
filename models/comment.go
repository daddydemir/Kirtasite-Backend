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
