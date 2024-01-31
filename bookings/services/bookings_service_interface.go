package services

import "gym-management/bookings/models/dtos"

type BookingsServiceInterface interface {
	GetBookings() *[]dtos.BookingCompleteDTO
	GetBooking(id int) (*dtos.BookingCompleteDTO, error)
	InsertNewBooking(newBooking *dtos.BookingDTO) (*dtos.BookingCompleteDTO, error)
	UpdateBooking(id int, updatedBooking *dtos.BookingDTO) (*dtos.BookingCompleteDTO, error, error)
	DeleteBooking(id int) (*dtos.BookingCompleteDTO, error, error)
	validateBooking(newBooking *dtos.BookingDTO) error
}