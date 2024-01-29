package dtos

import (
	"gym-management/bookings/models/dtos"
	"time"
)

type ClassScheduleWithBookingsDTO struct {
	Id        int
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Capacity  int
	Bookings []dtos.BookingCompleteDTO
}