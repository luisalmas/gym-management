package services

import (
	bookingsDtos "gym-management/bookings/models/dtos"
	bookingsRepo "gym-management/bookings/repositories"
	"gym-management/classes/models/dtos"
	"gym-management/classes/models/entities"
	"gym-management/classes/repositories"
	"time"
)

type ClassesServiceImpl struct {
	ClassesRepository repositories.ClassesRepository
	BookingsRepository bookingsRepo.BookingsRepository
}

func NewClassesService() *ClassesServiceImpl {
	return &ClassesServiceImpl{
		ClassesRepository: repositories.NewClassesRepository(),
		BookingsRepository: bookingsRepo.NewBookingsRepository(),
	}
}

func (service *ClassesServiceImpl) GetClasses() *[]dtos.ClassCompleteDTO {
	return service.ClassesRepository.GetClasses()
}

func (service *ClassesServiceImpl) InsertNewClass(classSchedule *dtos.ClassDTO) (*dtos.ClassCompleteDTO, error) {	
	classEntity := &entities.Class{}
	classScheduleEntity, err := classEntity.New(classSchedule)

	if err != nil{
		return nil, err
	}

	return service.ClassesRepository.InsertNewClass(classScheduleEntity), nil
}

func (service *ClassesServiceImpl) GetClass(id int) (*dtos.ClassCompleteDTO, error) {
	classEntity, err := service.ClassesRepository.GetClass(id)

	if err != nil{
		return nil, err
	}

	return classEntity.ToClassCompleteDTO(), err
}

func (service *ClassesServiceImpl) UpdateClass(id int, classSchedule *dtos.ClassDTO) (*dtos.ClassCompleteDTO, error) {
	currentClass, errGet := service.ClassesRepository.GetClass(id)

	if errGet != nil {
		return nil, errGet
	}

	classEntity := &entities.Class{}
	updatedClass, errUpdate := classEntity.New(classSchedule)

	if errUpdate != nil {
		return nil, errUpdate
	}

	updatedClass.ClassId = currentClass.ClassId

	//Cascade delete (for bookings outside the date range of the class)
	service.BookingsRepository.DeleteBookingsFromClass(updatedClass.ClassId, updatedClass.StartDate, updatedClass.EndDate)

	return service.ClassesRepository.UpdateClass(id, updatedClass), nil
}

func (service *ClassesServiceImpl) DeleteClass(id int) (*dtos.ClassCompleteDTO, error) {
	deletedClass, errorDelete := service.ClassesRepository.DeleteClass(id)

	if errorDelete != nil {
		return nil, errorDelete
	}

	//Cascade delete
	service.BookingsRepository.DeleteBookingsFromClass(deletedClass.ClassId, time.Time{}, time.Time{})

	return deletedClass, nil
}

func (service *ClassesServiceImpl) GetBookingsFromClass(id int, date time.Time) (*[]bookingsDtos.BookingCompleteDTO, error){
	_, errGet := service.ClassesRepository.GetClass(id)

	if errGet != nil {
		return nil, errGet
	}

	return service.BookingsRepository.GetBookingsFromClass(id, date), nil
}