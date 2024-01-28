package services

import (
	"gym-management/models/dtos"
	"gym-management/models/entities"
	"gym-management/repositories"
)

type ClassesService struct {
	ClassesRepository *repositories.ClassesRepository
}

func NewClassesService() *ClassesService {
	return &ClassesService{
		ClassesRepository: &repositories.ClassesRepository{},
	}
}

func (service *ClassesService) GetClassesSchedules() *[]dtos.ClassScheduleCompleteDTO {
	return service.ClassesRepository.GetClassesSchedules()
}

func (service *ClassesService) InsertNewClassSchedule(classSchedule *dtos.ClassScheduleDTO) (*dtos.ClassScheduleCompleteDTO, error) {
	classEntity := &entities.ClassSchedule{}
	classScheduleEntity, err := classEntity.New(classSchedule)

	if err != nil{
		return nil, err
	}

	return service.ClassesRepository.InsertNewClassSchedule(classScheduleEntity)
}

func (service *ClassesService) GetClassSchedule(id int) (*dtos.ClassScheduleCompleteDTO, error) {
	classEntity, err := service.ClassesRepository.GetClassSchedule(id)
	return classEntity.ToClassSheduleDTO(), err
}

func (service *ClassesService) UpdateClassSchedule(id int, classSchedule *dtos.ClassScheduleDTO) (*dtos.ClassScheduleCompleteDTO, error, error) {
	currentClass, errGet := service.ClassesRepository.GetClassSchedule(id)

	if errGet != nil {
		return nil, errGet, nil
	}

	classEntity := &entities.ClassSchedule{}
	updatedClass, errUpdate := classEntity.New(classSchedule)

	if errUpdate != nil {
		return nil, nil, errUpdate
	}

	updatedClass.Id = currentClass.Id

	return service.ClassesRepository.UpdateClassSchedule(id, updatedClass), nil, nil
}