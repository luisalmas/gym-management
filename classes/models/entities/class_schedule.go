package entities

import (
	"errors"
	"gym-management/classes/models/dtos"
	"time"
)

type ClassSchedule struct {
	ClassId		   int
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
		ClassId: 0,
		Name: classScheduleDTO.Name,
		StartDate: classScheduleDTO.StartDate,
		EndDate: classScheduleDTO.EndDate,
		Capacity: classScheduleDTO.Capacity,
	}, nil
}

func (class *ClassSchedule) ToClassSheduleDTO() (*dtos.ClassScheduleCompleteDTO) {
	return &dtos.ClassScheduleCompleteDTO{
		ClassId: class.ClassId,
		Name: class.Name,
		StartDate: class.StartDate,
		EndDate: class.EndDate,
		Capacity: class.Capacity,
	}
}

/*func (class *ClassSchedule) ToClassSheduleWithBookingsDTO(classBookings []bookingDtos.BookingCompleteDTO) (*dtos.ClassScheduleWithBookingsDTO) {
	return &dtos.ClassScheduleWithBookingsDTO{
		ClassId: class.ClassId,
		Name: class.Name,
		StartDate: class.StartDate,
		EndDate: class.EndDate,
		Capacity: class.Capacity,
		Bookings: classBookings,
	}
}*/