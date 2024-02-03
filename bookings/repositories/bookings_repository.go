package repositories

import (
	"gym-management/bookings/models/dtos"
	"gym-management/bookings/models/entities"
	"gym-management/bookings/models/errors"
	"time"
)

type BookingsRepositoryImpl struct {
	bookings []entities.Booking
}

func NewBookingsRepository() *BookingsRepositoryImpl{
	return &BookingsRepositoryImpl{
		bookings: []entities.Booking {
			{
				BookingId: 1,
				Name: "Peter",
				Date: time.Date(2024, time.January, 25,  0, 0, 0, 0, time.UTC),
				ClassId: 1,
			},
			{
				BookingId: 2,
				Name: "Samantha",
				Date: time.Date(2024, time.January, 26,  0, 0, 0, 0, time.UTC),
				ClassId: 1,
			},
		},
	}
}

func (repo *BookingsRepositoryImpl) GetBookings() *[]dtos.BookingCompleteDTO {
	bookingsDTO := []dtos.BookingCompleteDTO{}
	for _, booking := range repo.bookings {
		bookingsDTO = append(bookingsDTO, *booking.ToBookingCompleteDTO())
	}
	return &bookingsDTO
}

func (repo *BookingsRepositoryImpl) GetBookingsFromClass(classId int, date time.Time) *[]dtos.BookingCompleteDTO {
	bookingsDTO := []dtos.BookingCompleteDTO{}
	for _, booking := range repo.bookings {
		if booking.ClassId == classId {
			if date.IsZero() {
				bookingsDTO = append(bookingsDTO, *booking.ToBookingCompleteDTO())
			}else if booking.Date == date{
				bookingsDTO = append(bookingsDTO, *booking.ToBookingCompleteDTO())
			}
		}
	}
	return &bookingsDTO
}

func (repo *BookingsRepositoryImpl) GetBooking(id int) (*entities.Booking, error) {
	for index, booking := range repo.bookings {
		if booking.BookingId == id{
			return &repo.bookings[index], nil
		}
	}
	return nil, errors.NewBookingNotFoundError()
}

func (repo * BookingsRepositoryImpl) InsertNewBooking(newBooking *entities.Booking) *dtos.BookingCompleteDTO {
	newBooking.BookingId = len(repo.bookings) + 1
	repo.bookings = append(repo.bookings, *newBooking)
	return newBooking.ToBookingCompleteDTO()
}

func (repo * BookingsRepositoryImpl) UpdateBooking(id int, updatedBooking *entities.Booking) (*dtos.BookingCompleteDTO) {
	//Already done in service (simulate DB)
	currentBooking, _ := repo.GetBooking(id)

	currentBooking.Name = updatedBooking.Name
	currentBooking.Date = updatedBooking.Date

	return currentBooking.ToBookingCompleteDTO()
}

func (repo * BookingsRepositoryImpl) DeleteBooking(id int) (*dtos.BookingCompleteDTO, error) {
	for index, booking := range repo.bookings {
		if booking.BookingId == id{
			deletedBooking := booking
			repo.bookings = append(repo.bookings[:index], repo.bookings[index+1:]...)
			return deletedBooking.ToBookingCompleteDTO(), nil
		}
	}
	return nil, errors.NewBookingNotFoundError()
}

func (repo *BookingsRepositoryImpl) DeleteBookingsFromClass(classId int, dateStart time.Time, dateEnd time.Time) *[]dtos.BookingCompleteDTO{
	bookingsFromClass := repo.GetBookingsFromClass(classId ,time.Time{})
	deletedBookings := []dtos.BookingCompleteDTO{}
	useDates := !dateStart.IsZero() && !dateEnd.IsZero()
	for _, booking := range *bookingsFromClass{
		if useDates {
			if (booking.Date.Compare(dateStart) == -1  || booking.Date.Compare(dateEnd) == 1) {
				deletedBooking, _ := repo.DeleteBooking(booking.BookingId)
				deletedBookings = append(deletedBookings, *deletedBooking)
			}
			continue
		}

		deletedBooking, _ := repo.DeleteBooking(booking.BookingId)
		deletedBookings = append(deletedBookings, *deletedBooking)
	}

	return &deletedBookings
}