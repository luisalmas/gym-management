package repositories

import (
	"gym-management/bookings/models/dtos"
	"gym-management/bookings/models/entities"
	"time"
)

type BookingsRepository interface {
	GetBookings() *[]dtos.BookingCompleteDTO
	GetBooking(id int) (*entities.Booking, error)
	GetBookingsFromClass(classId int, date time.Time) *[]dtos.BookingCompleteDTO
	InsertNewBooking(newBooking *entities.Booking) *dtos.BookingCompleteDTO
	UpdateBooking(id int, updatedBooking *entities.Booking) *dtos.BookingCompleteDTO
	DeleteBooking(id int) (*dtos.BookingCompleteDTO, error)
	DeleteBookingsFromClass(classId int, dateStart time.Time, dateEnd time.Time) *[]dtos.BookingCompleteDTO
}