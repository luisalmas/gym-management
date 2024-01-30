package dtos

import "time"

type ClassDTO struct {
	Name       string 		`json:"name" binding:"required"`
	StartDate time.Time	`json:"startDate" binding:"required"`
	EndDate   time.Time	`json:"endDate" binding:"required"`
	Capacity   int			`json:"capacity" binding:"required"`
}

func (class *ClassDTO) HasValidDates() bool {
	return class.StartDate.Compare(class.EndDate) == 1
}