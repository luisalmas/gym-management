package entities

import (
	"gym-management/models/dtos"
	"time"
)

type Booking struct {
	Id	int
	Name       string
	Date time.Time
}

func (booking *Booking) ToBookingDTO() (*dtos.BookingCompleteDTO) {
	return &dtos.BookingCompleteDTO{
		Id: booking.Id,
		Name: booking.Name,
		Date: booking.Date,
	}
}