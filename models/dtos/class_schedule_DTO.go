package dtos

import "time"

type ClassScheduleDTO struct {
	Name       string 		`json:"name" binding:"required"`
	StartDate time.Time	`json:"start_date" binding:"required"`
	EndDate   time.Time	`json:"end_date" binding:"required"`
	Capacity   int			`json:"capacity" binding:"required"`
}

func (class *ClassScheduleDTO) HasValidDates() bool {
	return class.StartDate.Compare(class.EndDate) == 1
}