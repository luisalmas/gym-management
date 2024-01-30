package dtos

import "time"

type ClassCompleteDTO struct {
	ClassId        int
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Capacity  int
}

