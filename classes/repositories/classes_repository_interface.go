package repositories

import (
	"gym-management/classes/models/dtos"
	"gym-management/classes/models/entities"
)

type ClassesRepository interface {
	GetClasses() *[]dtos.ClassCompleteDTO
	InsertNewClass(classSchedule *entities.Class) (*dtos.ClassCompleteDTO)
	GetClass(id int) (*entities.Class, error)
	UpdateClass(id int, updatedClass *entities.Class) *dtos.ClassCompleteDTO
	DeleteClass(id int) (*dtos.ClassCompleteDTO, error)
}