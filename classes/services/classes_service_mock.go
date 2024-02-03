package services

import (
	bookingsDtos "gym-management/bookings/models/dtos"
	"gym-management/classes/models/dtos"
	"time"

	"github.com/stretchr/testify/mock"
)

type MockClassesService struct {
	mock.Mock
}

func (mockClassesService *MockClassesService)GetClasses() *[]dtos.ClassCompleteDTO{
	args := mockClassesService.Called()
	return args.Get(0).(*[]dtos.ClassCompleteDTO)
}
func (mockClassesService *MockClassesService)InsertNewClass(classSchedule *dtos.ClassDTO) (*dtos.ClassCompleteDTO, error){
	args := mockClassesService.Called()
	if args.Get(0) != nil && args.Get(1) != nil {
		return args.Get(0).(*dtos.ClassCompleteDTO), args.Get(1).(error)
	}

	if args.Get(1) == nil{
		return args.Get(0).(*dtos.ClassCompleteDTO), nil
	}

	if args.Get(0) == nil{
		return nil, args.Get(1).(error)
	}
	
	return nil, nil
}
func (mockClassesService *MockClassesService)GetClass(id int) (*dtos.ClassCompleteDTO, error){
	args := mockClassesService.Called()
	if args.Get(0) != nil && args.Get(1) != nil {
		return args.Get(0).(*dtos.ClassCompleteDTO), args.Get(1).(error)
	}

	if args.Get(1) == nil{
		return args.Get(0).(*dtos.ClassCompleteDTO), nil
	}

	if args.Get(0) == nil{
		return nil, args.Get(1).(error)
	}
	
	return nil, nil
}
func (mockClassesService *MockClassesService)UpdateClass(id int, classSchedule *dtos.ClassDTO) (*dtos.ClassCompleteDTO, error){
	args := mockClassesService.Called()
	if args.Get(0) != nil && args.Get(1) != nil {
		return args.Get(0).(*dtos.ClassCompleteDTO), args.Get(1).(error)
	}

	if args.Get(1) == nil{
		return args.Get(0).(*dtos.ClassCompleteDTO), nil
	}

	if args.Get(0) == nil{
		return nil, args.Get(1).(error)
	}
	
	return nil, nil
}
func (mockClassesService *MockClassesService)DeleteClass(id int) (*dtos.ClassCompleteDTO, error){
	args := mockClassesService.Called()
	if args.Get(0) != nil && args.Get(1) != nil {
		return args.Get(0).(*dtos.ClassCompleteDTO), args.Get(1).(error)
	}

	if args.Get(1) == nil{
		return args.Get(0).(*dtos.ClassCompleteDTO), nil
	}

	if args.Get(0) == nil{
		return nil, args.Get(1).(error)
	}
	
	return nil, nil
}
func (mockClassesService *MockClassesService)GetBookingsFromClass(id int, date time.Time) (*[]bookingsDtos.BookingCompleteDTO, error){
	args := mockClassesService.Called()
	if args.Get(0) != nil && args.Get(1) != nil {
		return args.Get(0).(*[]bookingsDtos.BookingCompleteDTO), args.Get(1).(error)
	}

	if args.Get(1) == nil{
		return args.Get(0).(*[]bookingsDtos.BookingCompleteDTO), nil
	}

	if args.Get(0) == nil{
		return nil, args.Get(1).(error)
	}
	
	return nil, nil
}