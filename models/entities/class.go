package entities

import "time"

type Class struct {
	Id	int
	Name       string
	Date time.Time
	Capacity   int
}