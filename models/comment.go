package models

import "time"

type Comment struct {
	Id           int
	CustomerId   int
	StationeryId int
	Content      string
	CreatedDate  time.Time
	Score        int
}
