package dtos

import "time"

type ClassScheduleCompleteDTO struct {
	Id        int
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Capacity  int
}

