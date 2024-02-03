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

func (class *Class) New(classDTO *dtos.ClassDTO) (*Class, error) {

	if classDTO.StartDate.Compare(classDTO.EndDate) == 1 {
		return nil, errors.New("Class: invalid dates")
	}

	if classDTO.Capacity < 1 {
		return nil, errors.New("Class: cannot have capacity less than 1")
	}

	return &Class{
		ClassId: 0,
		Name: classDTO.Name,
		StartDate: classDTO.StartDate,
		EndDate: classDTO.EndDate,
		Capacity: classDTO.Capacity,
	}, nil
}

func (class *Class) ToClassCompleteDTO() (*dtos.ClassCompleteDTO) {
	return &dtos.ClassCompleteDTO{
		ClassId: class.ClassId,
		Name: class.Name,
		StartDate: class.StartDate,
		EndDate: class.EndDate,
		Capacity: class.Capacity,
	}
}