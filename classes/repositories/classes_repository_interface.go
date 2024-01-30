package repositories

import (
	"gym-management/classes/models/dtos"
	"gym-management/classes/models/entities"
)

type ClassesRepositoryInterface interface {
	GetClassesSchedules() *[]dtos.ClassCompleteDTO
	InsertNewClassSchedule(classSchedule *entities.Class) (*dtos.ClassCompleteDTO, error)
	GetClassSchedule(id int) (*entities.Class, error)
	UpdateClassSchedule(id int, updatedClass *entities.Class) *dtos.ClassCompleteDTO
	DeleteClassSchedule(id int) (*dtos.ClassCompleteDTO, error)
}