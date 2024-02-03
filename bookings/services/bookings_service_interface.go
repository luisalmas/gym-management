package services

import (
	"gym-management/bookings/models/dtos"
)

type BookingsService interface {
	GetBookings() *[]dtos.BookingCompleteDTO
	GetBooking(id int) (*dtos.BookingCompleteDTO, error)
	InsertNewBooking(newBooking *dtos.BookingDTO) (*dtos.BookingCompleteDTO, error)
	UpdateBooking(id int, updatedBooking *dtos.BookingDTO) (*dtos.BookingCompleteDTO, error)
	DeleteBooking(id int) (*dtos.BookingCompleteDTO, error)
	validateBooking(newBooking *dtos.BookingDTO) error
}