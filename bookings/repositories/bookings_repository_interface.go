package repositories

import (
	"gym-management/bookings/models/dtos"
	"gym-management/bookings/models/entities"
)

type BookingsRepositoryInterface interface {
	GetBookings() *[]dtos.BookingCompleteDTO
	GetBooking(id int) (*entities.Booking, error)
	GetBookingsFromClass(classId int) *[]dtos.BookingCompleteDTO
	InsertNewBooking(newBooking *entities.Booking) (*dtos.BookingCompleteDTO, error)
	UpdateBooking(id int, updatedBooking *entities.Booking) *dtos.BookingCompleteDTO
	DeleteBooking(id int) (*dtos.BookingCompleteDTO, error)
}