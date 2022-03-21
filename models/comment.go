package models

import "time"

type Comment struct {
	Id           int       `json:"id"`
	UserId       int       `json:"user_id" validate:"required"`
	StationeryId int       `json:"stationery_id" validate:"required"`
	Content      string    `json:"content" validate:"required,max=499"`
	Date         time.Time `json:"date"`
	Score        float32   `json:"score" validate:"required,gt=0,lt=11"`
}
