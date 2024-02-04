package dtos

import (
	"time"
)

type BookingDTO struct {
	Name string `json:"name" binding:"required" error:"name is required"`
	Date time.Time `json:"date" binding:"required" error:"invalid date"`
	ClassId int `json:"classId" binding:"required" error:"classId is required"`
} 