package dtos

import "time"

type BookingDTO struct {
	Name string `json:"name" binding:"required"`
	Date time.Time `json:"date" binding:"required"`
	ClassId int `json:"classId" binding:"required"`
} 