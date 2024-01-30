package entities

import (
	"gym-management/bookings/models/dtos"
	"time"
)

type Booking struct {
	BookingId	int
	Name       string
	Date time.Time
	ClassId int
}

func (booking *Booking) ToBookingDTO() (*dtos.BookingCompleteDTO) {
	return &dtos.BookingCompleteDTO{
		BookingId: booking.BookingId,
		Name: booking.Name,
		Date: booking.Date,
		ClassId: booking.ClassId,
	}
}