package models

import "time"

type Order struct {
	Id           int
	FileId       int
	CustomerId   int
	StationeryId int
	StatusId     int
	TotalPrice   float32
	CreatedDate  time.Time
	DeliveryDate time.Time
}
