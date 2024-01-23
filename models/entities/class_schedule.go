package entities

import "time"

type ClassSchedule struct {
	Id		   int
	Name       string
	Start_date time.Time
	End_date   time.Time
	Capacity   int
}