package dtos

import "time"

type BookingCompleteDTO struct {
	BookingId   int
	Name string
	Date time.Time
	ClassId int
}