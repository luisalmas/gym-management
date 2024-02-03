package services

import (
	"gym-management/bookings/models/dtos"

	"github.com/stretchr/testify/mock"
)

type MockBookingsService struct {
	mock.Mock
}

func (mockBookingsService *MockBookingsService) GetBookings() *[]dtos.BookingCompleteDTO {
	args := mockBookingsService.Called()
	return args.Get(0).(*[]dtos.BookingCompleteDTO)
}

func (mockBookingsService *MockBookingsService) GetBooking(id int) (*dtos.BookingCompleteDTO, error) {
	args := mockBookingsService.Called(id)

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

func (mockBookingsService *MockBookingsService) InsertNewBooking(newBooking *dtos.BookingDTO) (*dtos.BookingCompleteDTO, error){
	args := mockBookingsService.Called(newBooking)

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

func (mockBookingsService *MockBookingsService) UpdateBooking(id int, updatedBooking *dtos.BookingDTO) (*dtos.BookingCompleteDTO, error){
	args := mockBookingsService.Called(id ,updatedBooking)
	
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

func (mockBookingsService *MockBookingsService) DeleteBooking(id int) (*dtos.BookingCompleteDTO, error){
	args := mockBookingsService.Called(id)
	
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

func (mockBookingsService *MockBookingsService) validateBooking(newBooking *dtos.BookingDTO) error {
	//Its not needed
	return nil
}