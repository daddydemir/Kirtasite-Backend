package models

import "time"

type Comment struct {
	Id           int       `json:"id"`
	UserId       int       `json:"user_id"`
	StationeryId int       `json:"stationery_id"`
	Content      string    `json:"content"`
	Date         time.Time `json:"date"`
	Score        float32   `json:"score"`
}
