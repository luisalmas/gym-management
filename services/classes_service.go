package services

import (
	"gym-management/models/entities"
	"gym-management/repositories"
)

type ClassesService struct {
	classesRepository *repositories.ClassesRepository
}

func NewClassesService() *ClassesService {
	return &ClassesService{
		classesRepository: &repositories.ClassesRepository{},
	}
}

func (serv *ClassesService) GetClassesSchedules() []entities.ClassSchedule {
	return serv.classesRepository.GetClassesSchedules()
}