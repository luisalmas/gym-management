package repositories

import (
	"errors"
	"gym-management/bookings/models/dtos"
	"gym-management/bookings/models/entities"
	"time"
)

var bookings = []entities.Booking {
	{
		Id: 1,
		Name: "Peter",
		Date: time.Date(2024, time.January, 25,  0, 0, 0, 0, time.UTC),
	},
	{
		Id: 2,
		Name: "Samantha",
		Date: time.Date(2024, time.January, 25,  0, 0, 0, 0, time.UTC),
	},
}

type BookingsRepository struct {
	//db connection
}

func NewBookingsRepository() *BookingsRepository{
	return &BookingsRepository{}
}

func (repo *BookingsRepository) GetBookings() *[]dtos.BookingCompleteDTO {
	bookingsDTO := []dtos.BookingCompleteDTO{}
	for _, booking := range bookings {
		bookingsDTO = append(bookingsDTO, *booking.ToBookingDTO())
	}
	return &bookingsDTO
}

func (repo *BookingsRepository) GetBooking(id int) (*entities.Booking, error) {
	for index, booking := range bookings {
		if booking.Id == id{
			return &bookings[index], nil
		}
	}
	return nil, errors.New("booking not found")
}

func (repo * BookingsRepository) InsertNewBooking(newBooking *entities.Booking) (*dtos.BookingCompleteDTO, error) {
	bookings = append(bookings, *newBooking)
	return newBooking.ToBookingDTO(), nil
}

func (repo * BookingsRepository) UpdateBooking(id int, updatedBooking *entities.Booking) (*dtos.BookingCompleteDTO) {
	//Already done in service (simulate DB)
	currentBooking, _ := repo.GetBooking(id)

	currentBooking.Name = updatedBooking.Name
	currentBooking.Date = updatedBooking.Date

	return currentBooking.ToBookingDTO()
}

func (repo * BookingsRepository) DeleteBooking(id int) (*dtos.BookingCompleteDTO, error) {
	for index, booking := range bookings {
		if booking.Id == id{
			deletedBooking := booking
			bookings = append(bookings[:index], bookings[index+1:]...)
			return deletedBooking.ToBookingDTO(), nil
		}
	}
	return nil, errors.New("no booking to delete")
}