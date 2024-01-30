package dtos

import "time"

type ClassScheduleCompleteDTO struct {
	ClassId        int
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Capacity  int
}

