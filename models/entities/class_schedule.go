package entities

import (
	"errors"
	"gym-management/models/dtos"
	"time"
)

type ClassSchedule struct {
	Id		   int
	Name       string
	Start_date time.Time
	End_date   time.Time
	Capacity   int
}

func (class *ClassSchedule) New(classScheduleDTO *dtos.ClassScheduleDTO) (*ClassSchedule, error) {

	if classScheduleDTO.Start_date.Compare(classScheduleDTO.End_date) == 1 {
		return nil, errors.New("ClassSchedule")
	}

	return &ClassSchedule{
		Id: 0,
		Name: classScheduleDTO.Name,
		Start_date: classScheduleDTO.Start_date,
		End_date: classScheduleDTO.End_date,
		Capacity: classScheduleDTO.Capacity,
	}, nil
}