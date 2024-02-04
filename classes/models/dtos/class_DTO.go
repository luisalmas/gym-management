package dtos

import "time"

type ClassDTO struct {
	Name       string 		`json:"name" binding:"required" required:"name is required"`
	StartDate time.Time	`json:"startDate" binding:"required" required:"invalid start date"`
	EndDate   time.Time	`json:"endDate" binding:"required,gtfield=StartDate" required:"invalid end date" gtfield:"endDate cannot be greater than the startDate"` 
	Capacity   int			`json:"capacity" binding:"required,gt=0" required:"capacity must be greater than 0"`
}