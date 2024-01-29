package services

import (
	"gym-management/classes/models/dtos"
	"gym-management/classes/models/entities"
	"gym-management/classes/repositories"
)

type ClassesService struct {
	ClassesRepository repositories.ClassesRepositoryInterface
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

func (service *ClassesService) DeleteClassSchedule(id int) (*dtos.ClassScheduleCompleteDTO, error, error) {
	currentClass, errGet := service.ClassesRepository.GetClassSchedule(id)

	if errGet != nil {
		return nil, errGet, nil
	}

	deletedClass, errorDelete := service.ClassesRepository.DeleteClassSchedule(currentClass.Id)

	if errorDelete != nil {
		return nil, nil, errorDelete
	}

	return deletedClass, nil, nil
}