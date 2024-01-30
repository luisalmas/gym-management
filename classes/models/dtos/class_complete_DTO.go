package dtos

import "time"

type ClassCompleteDTO struct {
	ClassId        int
	Name       string 		`json:"name"`
	StartDate time.Time	`json:"startDate"`
	EndDate   time.Time	`json:"endDate"`
	Capacity   int			`json:"capacity"`
}

