package dtos

import "time"

type ClassDTO struct {
	Name       string
	Start_date time.Time
	End_date   time.Time
	Capacity   int
}