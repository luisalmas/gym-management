package services

import (
	bookingsDtos "gym-management/bookings/models/dtos"
	"gym-management/classes/models/dtos"
	"time"
)

type ClassesService interface {
	GetClasses() *[]dtos.ClassCompleteDTO
	InsertNewClass(classSchedule *dtos.ClassDTO) (*dtos.ClassCompleteDTO, error)
	GetClass(id int) (*dtos.ClassCompleteDTO, error)
	UpdateClass(id int, classSchedule *dtos.ClassDTO) (*dtos.ClassCompleteDTO, error)
	DeleteClass(id int) (*dtos.ClassCompleteDTO, error)
	GetBookingsFromClass(id int, date time.Time) (*[]bookingsDtos.BookingCompleteDTO, error)
}