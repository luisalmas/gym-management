package dtos

import "time"

type BookingCompleteDTO struct {
	BookingId   int `json:"bookingId" binding:"required"`
	Name string `json:"name" binding:"required"`
	Date time.Time `json:"date" binding:"required"`
	ClassId int `json:"classId" binding:"required"`
}