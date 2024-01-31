package services

import (
	bookingsDtos "gym-management/bookings/models/dtos"
	bookingsRepo "gym-management/bookings/repositories"
	"gym-management/classes/models/dtos"
	"gym-management/classes/models/entities"
	"gym-management/classes/repositories"
	"time"
)

type ClassesService struct {
	ClassesRepository repositories.ClassesRepositoryInterface
	BookingsRepository bookingsRepo.BookingsRepositoryInterface
}

func NewClassesService() *ClassesService {
	return &ClassesService{
		ClassesRepository: &repositories.ClassesRepository{},
		BookingsRepository: &bookingsRepo.BookingsRepository{},
	}
}

func (service *ClassesService) GetClassesSchedules() *[]dtos.ClassCompleteDTO {
	return service.ClassesRepository.GetClassesSchedules()
}

func (service *ClassesService) InsertNewClassSchedule(classSchedule *dtos.ClassDTO) (*dtos.ClassCompleteDTO, error) {
	classEntity := &entities.Class{}
	classScheduleEntity, err := classEntity.New(classSchedule)

	if err != nil{
		return nil, err
	}

	return service.ClassesRepository.InsertNewClassSchedule(classScheduleEntity)
}

func (service *ClassesService) GetClassSchedule(id int) (*dtos.ClassCompleteDTO, error) {
	classEntity, err := service.ClassesRepository.GetClassSchedule(id)
	return classEntity.ToClassSheduleDTO(), err
}

func (service *ClassesService) UpdateClassSchedule(id int, classSchedule *dtos.ClassDTO) (*dtos.ClassCompleteDTO, error, error) {
	currentClass, errGet := service.ClassesRepository.GetClassSchedule(id)

	if errGet != nil {
		return nil, errGet, nil
	}

	classEntity := &entities.Class{}
	updatedClass, errUpdate := classEntity.New(classSchedule)

	if errUpdate != nil {
		return nil, nil, errUpdate
	}

	updatedClass.ClassId = currentClass.ClassId

	return service.ClassesRepository.UpdateClassSchedule(id, updatedClass), nil, nil
}

func (service *ClassesService) DeleteClassSchedule(id int) (*dtos.ClassCompleteDTO, error, error) {
	currentClass, errGet := service.ClassesRepository.GetClassSchedule(id)

	if errGet != nil {
		return nil, errGet, nil
	}

	deletedClass, errorDelete := service.ClassesRepository.DeleteClassSchedule(currentClass.ClassId)

	if errorDelete != nil {
		return nil, nil, errorDelete
	}

	service.BookingsRepository.DeleteBookingsFromClass(deletedClass.ClassId, time.Time{})

	return deletedClass, nil, nil
}

func (service *ClassesService) GetBookingsFromClass(id int, date time.Time) (*[]bookingsDtos.BookingCompleteDTO, error){
	_, errGet := service.ClassesRepository.GetClassSchedule(id)

	if errGet != nil {
		return nil, errGet
	}

	return service.BookingsRepository.GetBookingsFromClass(id, date), nil
}