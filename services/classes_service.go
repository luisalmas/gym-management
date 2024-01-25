package services

import (
	"fmt"
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
	entity := &entities.ClassSchedule{}
	classScheduleEntity, err := entity.New(classSchedule)

	if err != nil{
		return nil, err
	}

	return service.classesRepository.InsertNewClassSchedule(classScheduleEntity)
}

func (service *ClassesService) GetClassSchedule(id int) (*entities.ClassSchedule, error) {
	return service.classesRepository.GetClassSchedule(id)
}

func (service *ClassesService) UpdateClassSchedule(id int, classSchedule *dtos.ClassScheduleDTO) (*entities.ClassSchedule, error) {
	_, err := service.classesRepository.GetClassSchedule(id)

	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	return service.classesRepository.UpdateClassSchedule(id, classSchedule)
}