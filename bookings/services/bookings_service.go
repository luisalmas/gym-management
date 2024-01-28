package services

import (
	"gym-management/bookings/models/dtos"
	"gym-management/bookings/models/entities"
	"gym-management/bookings/repositories"
)

type BookingsService struct {
	BookingsRepository *repositories.BookingsRepository
}

func NewBookingsService() *BookingsService {
	return &BookingsService{
		BookingsRepository: &repositories.BookingsRepository{},
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

	//Perform booking validations...
	//....

	return service.BookingsRepository.InsertNewBooking(&entities.Booking{Name: newBooking.Name, Date: newBooking.Date})
}

func (service *BookingsService) UpdateBooking(id int, updatedBooking *dtos.BookingDTO) (*dtos.BookingCompleteDTO, error, error){
	currentBooking, errGet := service.BookingsRepository.GetBooking(id)
	
	if errGet != nil {
		return nil, errGet, nil
	}

	bookingEntity := &entities.Booking{
		Id: currentBooking.Id,
		Name: updatedBooking.Name,
		Date: updatedBooking.Date,
	}

	//Perform booking validations...
	//....

	return service.BookingsRepository.UpdateBooking(id, bookingEntity), nil, nil
}