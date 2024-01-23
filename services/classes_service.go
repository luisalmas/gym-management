package services

import (
	"gym-management/models/dtos"
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

func (service *ClassesService) GetClassesSchedules() *[]entities.ClassSchedule {
	return service.classesRepository.GetClassesSchedules()
}

func (service *ClassesService) InsertNewClassSchedule(classSchedule *dtos.ClassScheduleDTO) (*entities.ClassSchedule, error) {
	return service.classesRepository.InsertNewClassSchedule(classSchedule)
}