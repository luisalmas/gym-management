package dtos

import "time"

type ClassDTO struct {
	Name       string 		`json:"name" binding:"required" error:"name is required"`
	StartDate time.Time	`json:"startDate" binding:"required" error:"invalid start date"`
	EndDate   time.Time	`json:"endDate" binding:"required,gtfield=StartDate" error:"invalid end date"`
	Capacity   int			`json:"capacity" binding:"required,gt=0" error:"capacity must be greater than 0"`
}