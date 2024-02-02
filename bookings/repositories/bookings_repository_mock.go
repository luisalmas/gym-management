package repositories

import (
	"gym-management/bookings/models/dtos"
	"gym-management/bookings/models/entities"
	"time"

	"github.com/stretchr/testify/mock"
)

type MockBookingsRepository struct {
	mock.Mock
}

func (mockBookingRepo *MockBookingsRepository) GetBookings() *[]dtos.BookingCompleteDTO {
	args := mockBookingRepo.Called()
	return args.Get(0).(*[]dtos.BookingCompleteDTO)
}

func (mockBookingRepo *MockBookingsRepository) GetBookingsFromClass(classId int, date time.Time) *[]dtos.BookingCompleteDTO {
	args := mockBookingRepo.Called(classId, date)
	return args.Get(0).(*[]dtos.BookingCompleteDTO)
}

func (mockBookingRepo *MockBookingsRepository) GetBooking(id int) (*entities.Booking, error) {
	args := mockBookingRepo.Called(id)

	if args.Get(0) != nil && args.Get(1) != nil {
		return args.Get(0).(*entities.Booking), args.Get(1).(error)
	}

	if args.Get(1) == nil{
		return args.Get(0).(*entities.Booking), nil
	}

	if args.Get(0) == nil{
		return nil, args.Get(1).(error)
	}

	return nil, nil
}

func (mockBookingRepo *MockBookingsRepository) InsertNewBooking(newBooking *entities.Booking) *dtos.BookingCompleteDTO {
	args := mockBookingRepo.Called(newBooking)
	return args.Get(0).(*dtos.BookingCompleteDTO)
}

func (mockBookingRepo *MockBookingsRepository) UpdateBooking(id int, updatedBooking *entities.Booking) *dtos.BookingCompleteDTO {
	args := mockBookingRepo.Called(id, updatedBooking)
	return args.Get(0).(*dtos.BookingCompleteDTO)
}

func (mockBookingRepo *MockBookingsRepository) DeleteBooking(id int) (*dtos.BookingCompleteDTO, error) {
	args := mockBookingRepo.Called(id)
	if args.Get(0) != nil && args.Get(1) != nil {
		return args.Get(0).(*dtos.BookingCompleteDTO), args.Get(1).(error)
	}

	if args.Get(1) == nil{
		return args.Get(0).(*dtos.BookingCompleteDTO), nil
	}

	if args.Get(0) == nil{
		return nil, args.Get(1).(error)
	}

	return nil, nil
}

func (mockBookingRepo *MockBookingsRepository) DeleteBookingsFromClass(classId int, dateStart time.Time, dateEnd time.Time) *[]dtos.BookingCompleteDTO {
	args := mockBookingRepo.Called(classId, dateStart, dateEnd)
	return args.Get(0).(*[]dtos.BookingCompleteDTO)
}