package entities

import (
	"errors"
	"gym-management/classes/models/dtos"
	"time"
)

type Class struct {
	ClassId		   int
	Name       string
	StartDate time.Time
	EndDate   time.Time
	Capacity   int
}

func (class *Class) New(classScheduleDTO *dtos.ClassDTO) (*Class, error) {

	if classScheduleDTO.StartDate.Compare(classScheduleDTO.EndDate) == 1 {
		return nil, errors.New("ClassSchedule: invalid dates")
	}

	return &Class{
		ClassId: 0,
		Name: classScheduleDTO.Name,
		StartDate: classScheduleDTO.StartDate,
		EndDate: classScheduleDTO.EndDate,
		Capacity: classScheduleDTO.Capacity,
	}, nil
}

func (class *Class) ToClassSheduleDTO() (*dtos.ClassCompleteDTO) {
	return &dtos.ClassCompleteDTO{
		ClassId: class.ClassId,
		Name: class.Name,
		StartDate: class.StartDate,
		EndDate: class.EndDate,
		Capacity: class.Capacity,
	}
}