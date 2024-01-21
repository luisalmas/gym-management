package entities

import "time"

type Booking struct {
	Id	int
	Name       string
	Date time.Time
	ClassId	int
}