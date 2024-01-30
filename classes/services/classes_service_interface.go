package services

import (
	bookingsDtos "gym-management/bookings/models/dtos"
	"gym-management/classes/models/dtos"
	"time"
)

type ClassesServiceInterface interface {
	GetClassesSchedules() *[]dtos.ClassScheduleCompleteDTO
	InsertNewClassSchedule(classSchedule *dtos.ClassScheduleDTO) (*dtos.ClassScheduleCompleteDTO, error)
	GetClassSchedule(id int) (*dtos.ClassScheduleCompleteDTO, error)
	UpdateClassSchedule(id int, classSchedule *dtos.ClassScheduleDTO) (*dtos.ClassScheduleCompleteDTO, error, error)
	DeleteClassSchedule(id int) (*dtos.ClassScheduleCompleteDTO, error, error)
	GetBookingsFromClass(id int, date time.Time) (*[]bookingsDtos.BookingCompleteDTO, error)
}