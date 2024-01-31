package repositories

import (
	"errors"
	"gym-management/bookings/models/dtos"
	"gym-management/bookings/models/entities"
	"time"
)

var bookings = []entities.Booking {
	{
		BookingId: 1,
		Name: "Peter",
		Date: time.Date(2024, time.January, 25,  0, 0, 0, 0, time.UTC),
		ClassId: 1,
	},
	{
		BookingId: 2,
		Name: "Samantha",
		Date: time.Date(2024, time.January, 25,  0, 0, 0, 0, time.UTC),
		ClassId: 1,
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

func (repo *BookingsRepository) GetBookingsFromClass(classId int, date time.Time) *[]dtos.BookingCompleteDTO {
	bookingsDTO := []dtos.BookingCompleteDTO{}
	for _, booking := range bookings {
		if booking.ClassId == classId {
			if date.IsZero() {
				bookingsDTO = append(bookingsDTO, *booking.ToBookingDTO())
			}else if booking.Date == date{
				bookingsDTO = append(bookingsDTO, *booking.ToBookingDTO())
			}
		}
	}
	return &bookingsDTO
}

func (repo *BookingsRepository) GetBooking(id int) (*entities.Booking, error) {
	for index, booking := range bookings {
		if booking.BookingId == id{
			return &bookings[index], nil
		}
	}
	return nil, errors.New("booking not found")
}

func (repo * BookingsRepository) InsertNewBooking(newBooking *entities.Booking) (*dtos.BookingCompleteDTO, error) {
	newBooking.BookingId = len(bookings) + 1
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
		if booking.BookingId == id{
			deletedBooking := booking
			bookings = append(bookings[:index], bookings[index+1:]...)
			return deletedBooking.ToBookingDTO(), nil
		}
	}
	return nil, errors.New("no booking to delete")
}

func (repo *BookingsRepository) DeleteBookingsFromClass(classId int, dateStart time.Time, dateEnd time.Time) *[]dtos.BookingCompleteDTO{
	bookingsFromClass := repo.GetBookingsFromClass(classId ,time.Time{})
	var deletedBookings []dtos.BookingCompleteDTO

	for _, booking := range *bookingsFromClass{
		if dateStart.IsZero() && dateEnd.IsZero() {
			deletedBooking, _ := repo.DeleteBooking(booking.BookingId)
			deletedBookings = append(deletedBookings, *deletedBooking)
		} else if booking.Date.Compare(dateStart) == -1  || booking.Date.Compare(dateEnd) == 1 {
			deletedBooking, _ := repo.DeleteBooking(booking.BookingId)
			deletedBookings = append(deletedBookings, *deletedBooking)
		}
	}

	return &deletedBookings
}