package services

import (
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
	_, errGetClass := service.ClassesRepository.GetClassSchedule(newBooking.ClassId)

	if errGetClass != nil {
		return nil, errGetClass
	}

	return service.BookingsRepository.InsertNewBooking(&entities.Booking{Name: newBooking.Name, Date: newBooking.Date, ClassId: newBooking.ClassId})
}

func (service *BookingsService) UpdateBooking(id int, updatedBooking *dtos.BookingDTO) (*dtos.BookingCompleteDTO, error, error){
	currentBooking, errGet := service.BookingsRepository.GetBooking(id)
	
	if errGet != nil {
		return nil, errGet, nil
	}

	bookingEntity := &entities.Booking{
		BookingId: currentBooking.BookingId,
		Name: updatedBooking.Name,
		Date: updatedBooking.Date,
	}

	//Perform booking validations...
	//....

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