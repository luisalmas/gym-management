package entities

import (
	"errors"
	bookingDtos "gym-management/bookings/models/dtos"
	"gym-management/classes/models/dtos"
	"time"
)

type ClassSchedule struct {
	Id		   int
	Name       string
	StartDate time.Time
	EndDate   time.Time
	Capacity   int
}

func (class *ClassSchedule) New(classScheduleDTO *dtos.ClassScheduleDTO) (*ClassSchedule, error) {

	if classScheduleDTO.StartDate.Compare(classScheduleDTO.EndDate) == 1 {
		return nil, errors.New("ClassSchedule: invalid dates")
	}

	return &ClassSchedule{
		Id: 0,
		Name: classScheduleDTO.Name,
		StartDate: classScheduleDTO.StartDate,
		EndDate: classScheduleDTO.EndDate,
		Capacity: classScheduleDTO.Capacity,
	}, nil
}

func (class *ClassSchedule) ToClassSheduleDTO() (*dtos.ClassScheduleCompleteDTO) {
	return &dtos.ClassScheduleCompleteDTO{
		Id: class.Id,
		Name: class.Name,
		StartDate: class.StartDate,
		EndDate: class.EndDate,
		Capacity: class.Capacity,
	}
}

func (class *ClassSchedule) ToClassSheduleWithBookingsDTO(classBookings []bookingDtos.BookingCompleteDTO) (*dtos.ClassScheduleWithBookingsDTO) {
	return &dtos.ClassScheduleWithBookingsDTO{
		Id: class.Id,
		Name: class.Name,
		StartDate: class.StartDate,
		EndDate: class.EndDate,
		Capacity: class.Capacity,
		Bookings: classBookings,
	}
}