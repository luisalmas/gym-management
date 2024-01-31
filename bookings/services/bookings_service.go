package services

import (
	"errors"
	"gym-management/bookings/models/dtos"
	"gym-management/bookings/models/entities"
	"gym-management/bookings/repositories"
	classesRepo "gym-management/classes/repositories"
)

type BookingsService struct {
	BookingsRepository repositories.BookingsRepositoryInterface
	ClassesRepository classesRepo.ClassesRepositoryInterface
}

func NewBookingsService() *BookingsService {
	return &BookingsService{
		BookingsRepository: repositories.NewBookingsRepository(),
		ClassesRepository: &classesRepo.ClassesRepository{},
	}
}

func (service *BookingsService) GetBookings() *[]dtos.BookingCompleteDTO {
	return service.BookingsRepository.GetBookings()
}

func (service *BookingsService) GetBooking(id int) (*dtos.BookingCompleteDTO, error) {
	bookingEntity, err := service.BookingsRepository.GetBooking(id)
	return bookingEntity.ToBookingDTO(), err
}

func (service *BookingsService) InsertNewBooking(newBooking *dtos.BookingDTO) (*dtos.BookingCompleteDTO, error){

	if err := service.validateBooking(newBooking); err != nil {
		return nil, err
	}

	return service.BookingsRepository.InsertNewBooking(&entities.Booking{Name: newBooking.Name, Date: newBooking.Date, ClassId: newBooking.ClassId})
}

func (service *BookingsService) UpdateBooking(id int, updatedBooking *dtos.BookingDTO) (*dtos.BookingCompleteDTO, error, error){
	currentBooking, errGet := service.BookingsRepository.GetBooking(id)
	
	if errGet != nil {
		return nil, errGet, nil
	}

	if err := service.validateBooking(updatedBooking); err != nil {
		return nil, nil, err
	}

	bookingEntity := &entities.Booking{
		BookingId: currentBooking.BookingId,
		Name: updatedBooking.Name,
		Date: updatedBooking.Date,
		ClassId: updatedBooking.ClassId,
	}

	return service.BookingsRepository.UpdateBooking(id, bookingEntity), nil, nil
}

func (service *BookingsService) DeleteBooking(id int) (*dtos.BookingCompleteDTO, error, error){
	currentBooking, errGet := service.BookingsRepository.GetBooking(id)
	
	if errGet != nil {
		return nil, errGet, nil
	}

	deletedBooking, errDelete := service.BookingsRepository.DeleteBooking(currentBooking.BookingId)

	if errGet != nil {
		return nil, nil, errDelete
	}

	return deletedBooking, nil, nil
}

func (service *BookingsService) validateBooking(newBooking *dtos.BookingDTO) error {
	class, errGetClass := service.ClassesRepository.GetClassSchedule(newBooking.ClassId)

	if errGetClass != nil {
		return errGetClass
	}

	if newBooking.Date.Compare(class.StartDate) == -1  || newBooking.Date.Compare(class.EndDate) == 1 {
		return errors.New("booking date does not correspond to the specified class")
	}
	return nil
}