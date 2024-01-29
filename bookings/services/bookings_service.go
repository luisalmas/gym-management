package services

import (
	"gym-management/bookings/models/dtos"
	"gym-management/bookings/models/entities"
	"gym-management/bookings/repositories"
	"gym-management/classes/services"
)

type BookingsService struct {
	BookingsRepository repositories.BookingsRepositoryInterface
	ClassesService services.ClassesServiceInterface
}

func NewBookingsService() *BookingsService {
	return &BookingsService{
		BookingsRepository: repositories.NewBookingsRepository(),
		ClassesService: services.NewClassesService(),
	}
}

func (service *BookingsService) GetBookings() *[]dtos.BookingCompleteDTO {
	return service.BookingsRepository.GetBookings()
}

func (service *BookingsService) GetBookingsFromClass(classId int) (*[]dtos.BookingCompleteDTO, error) {
	_, errGetClass := service.ClassesService.GetClassSchedule(classId)

	if errGetClass != nil {
		return nil, errGetClass
	}

	return service.BookingsRepository.GetBookingsFromClass(classId), nil
}

func (service *BookingsService) GetBooking(id int) (*dtos.BookingCompleteDTO, error) {
	bookingEntity, err := service.BookingsRepository.GetBooking(id)
	return bookingEntity.ToBookingDTO(), err
}

func (service *BookingsService) InsertNewBooking(newBooking *dtos.BookingDTO) (*dtos.BookingCompleteDTO, error){
	_, errGetClass := service.ClassesService.GetClassSchedule(newBooking.ClassId)

	if errGetClass != nil {
		return nil, errGetClass
	}

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

func (service *BookingsService) DeleteBooking(id int) (*dtos.BookingCompleteDTO, error, error){
	currentBooking, errGet := service.BookingsRepository.GetBooking(id)
	
	if errGet != nil {
		return nil, errGet, nil
	}

	deletedBooking, errDelete := service.BookingsRepository.DeleteBooking(currentBooking.Id)

	if errGet != nil {
		return nil, nil, errDelete
	}

	return deletedBooking, nil, nil
}