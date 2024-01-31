package services

import (
	bookingsDtos "gym-management/bookings/models/dtos"
	"gym-management/classes/models/dtos"
	"time"
)

type ClassesService interface {
	GetClassesSchedules() *[]dtos.ClassCompleteDTO
	InsertNewClassSchedule(classSchedule *dtos.ClassDTO) (*dtos.ClassCompleteDTO, error)
	GetClassSchedule(id int) (*dtos.ClassCompleteDTO, error)
	UpdateClassSchedule(id int, classSchedule *dtos.ClassDTO) (*dtos.ClassCompleteDTO, error, error)
	DeleteClassSchedule(id int) (*dtos.ClassCompleteDTO, error, error)
	GetBookingsFromClass(id int, date time.Time) (*[]bookingsDtos.BookingCompleteDTO, error)
}