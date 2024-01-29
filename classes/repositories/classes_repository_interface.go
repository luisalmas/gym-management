package repositories

import (
	"gym-management/classes/models/dtos"
	"gym-management/classes/models/entities"
)

type ClassesRepositoryInterface interface {
	GetClassesSchedules() *[]dtos.ClassScheduleCompleteDTO
	InsertNewClassSchedule(classSchedule *entities.ClassSchedule) (*dtos.ClassScheduleCompleteDTO, error)
	GetClassSchedule(id int) (*entities.ClassSchedule, error)
	UpdateClassSchedule(id int, updatedClass *entities.ClassSchedule) *dtos.ClassScheduleCompleteDTO
	DeleteClassSchedule(id int) (*dtos.ClassScheduleCompleteDTO, error)
}