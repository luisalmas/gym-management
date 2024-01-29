package dtos

import "time"

type BookingCompleteDTO struct {
	Id   int
	Name string
	Date time.Time
	ClassId int
}