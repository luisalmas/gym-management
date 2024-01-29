package entities

import (
	"gym-management/bookings/models/dtos"
	"time"
)

type Booking struct {
	Id	int
	Name       string
	Date time.Time
	ClassId int
}

func (booking *Booking) ToBookingDTO() (*dtos.BookingCompleteDTO) {
	return &dtos.BookingCompleteDTO{
		Id: booking.Id,
		Name: booking.Name,
		Date: booking.Date,
		ClassId: booking.ClassId,
	}
}