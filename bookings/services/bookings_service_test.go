package services

import (
	"gym-management/bookings/models/entities"
	"gym-management/bookings/repositories"
	classesRepo "gym-management/classes/repositories"
	"testing"
	"time"
)

var firstbooking = entities.Booking{
	BookingId: 1,
	Name:      "Peter",
	Date:      time.Date(2024, time.January, 25, 0, 0, 0, 0, time.UTC),
	ClassId:   1,
}

var secondbooking = entities.Booking{
	BookingId: 2,
	Name:      "Samantha",
	Date:      time.Date(2024, time.January, 26, 0, 0, 0, 0, time.UTC),
	ClassId:   1,
}

func TestGetBookings(t *testing.T) {
	bookingsService := NewBookingsService()
	bookingsService.BookingsRepository = new(repositories.MockBookingsRepository)
	bookingsService.ClassesRepository = new(classesRepo.MockClassesRepository)
}