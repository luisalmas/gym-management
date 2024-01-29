package services

import "gym-management/classes/models/dtos"

type ClassesServiceInterface interface {
	GetClassesSchedules() *[]dtos.ClassScheduleWithBookingsDTO
	InsertNewClassSchedule(classSchedule *dtos.ClassScheduleDTO) (*dtos.ClassScheduleCompleteDTO, error)
	GetClassSchedule(id int) (*dtos.ClassScheduleCompleteDTO, error)
	UpdateClassSchedule(id int, classSchedule *dtos.ClassScheduleDTO) (*dtos.ClassScheduleCompleteDTO, error, error)
	DeleteClassSchedule(id int) (*dtos.ClassScheduleCompleteDTO, error, error)
}