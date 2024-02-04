package dtos

import (
	"time"
)

type BookingDTO struct {
	Name string `json:"name" binding:"required" required:"name is required"`
	Date time.Time `json:"date" binding:"required" required:"invalid date"`
	ClassId int `json:"classId" binding:"required" required:"classId is required"`
} 