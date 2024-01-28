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
	classEntity := &entities.ClassSchedule{}
	classScheduleEntity, err := classEntity.New(classSchedule)

	if err != nil{
		return nil, err
	}

	return service.classesRepository.InsertNewClassSchedule(classScheduleEntity)
}

func (service *ClassesService) GetClassSchedule(id int) (*entities.ClassSchedule, error) {
	return service.classesRepository.GetClassSchedule(id)
}

func (service *ClassesService) UpdateClassSchedule(id int, classSchedule *dtos.ClassScheduleDTO) (*entities.ClassSchedule, error, error) {
	currentClass, errGet := service.classesRepository.GetClassSchedule(id)

	if errGet != nil {
		return nil, errGet, nil
	}

	classEntity := &entities.ClassSchedule{}
	updatedClass, errUpdate := classEntity.New(classSchedule)

	if errUpdate != nil {
		return nil, nil, errUpdate
	}

	updatedClass.Id = currentClass.Id

	return service.classesRepository.UpdateClassSchedule(id, updatedClass), nil, nil
}