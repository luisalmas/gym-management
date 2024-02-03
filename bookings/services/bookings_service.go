package services

import (
	"gym-management/bookings/models/dtos"
	"gym-management/bookings/models/entities"
	"gym-management/bookings/models/errors"
	"gym-management/bookings/repositories"
	classesRepo "gym-management/classes/repositories"
)

type BookingsServiceImpl struct {
	BookingsRepository repositories.BookingsRepository
	ClassesRepository classesRepo.ClassesRepository
}

func NewBookingsService() *BookingsServiceImpl {
	return &BookingsServiceImpl{
		BookingsRepository: repositories.NewBookingsRepository(),
		ClassesRepository: classesRepo.NewClassesRepository(),
	}
}

func (service *BookingsServiceImpl) GetBookings() *[]dtos.BookingCompleteDTO {
	return service.BookingsRepository.GetBookings()
}

func (service *BookingsServiceImpl) GetBooking(id int) (*dtos.BookingCompleteDTO, error) {
	bookingEntity, err := service.BookingsRepository.GetBooking(id)

	if err != nil {
		return nil, err
	}

	return bookingEntity.ToBookingDTO(), err
}

func (service *BookingsServiceImpl) InsertNewBooking(newBooking *dtos.BookingDTO) (*dtos.BookingCompleteDTO, error){

	if err := service.validateBooking(newBooking); err != nil {
		return nil, err
	}

	return service.BookingsRepository.InsertNewBooking(&entities.Booking{
		Name: newBooking.Name,
		Date: newBooking.Date,
		ClassId: newBooking.ClassId,
		}), nil
}

func (service *BookingsServiceImpl) UpdateBooking(id int, updatedBooking *dtos.BookingDTO) (*dtos.BookingCompleteDTO, error){
	currentBooking, errGet := service.BookingsRepository.GetBooking(id)
	
	if errGet != nil {
		return nil, errGet
	}

	if err := service.validateBooking(updatedBooking); err != nil {
		return nil, err
	}

	bookingEntity := &entities.Booking{
		BookingId: currentBooking.BookingId,
		Name: updatedBooking.Name,
		Date: updatedBooking.Date,
		ClassId: updatedBooking.ClassId,
	}

	return service.BookingsRepository.UpdateBooking(id, bookingEntity), nil
}

func (service *BookingsServiceImpl) DeleteBooking(id int) (*dtos.BookingCompleteDTO, error){
	deletedBooking, errDelete := service.BookingsRepository.DeleteBooking(id)

	if errDelete != nil {
		return nil, errDelete
	}

	return deletedBooking, nil
}

func (service *BookingsServiceImpl) validateBooking(newBooking *dtos.BookingDTO) error {
	class, errGetClass := service.ClassesRepository.GetClassSchedule(newBooking.ClassId)

	if errGetClass != nil {
		return errGetClass
	}

	if newBooking.Date.Compare(class.StartDate) == -1  || newBooking.Date.Compare(class.EndDate) == 1 {
		return errors.NewBookingDateInvalid()
	}
	return nil
}