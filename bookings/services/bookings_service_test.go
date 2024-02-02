package services

import (
	"errors"
	"fmt"
	"gym-management/bookings/models/dtos"
	"gym-management/bookings/models/entities"
	"gym-management/bookings/repositories"

	classesEntities "gym-management/classes/models/entities"
	classesRepo "gym-management/classes/repositories"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
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

var class = classesEntities.Class{
		ClassId: 1,
		Name: "Class1",
		StartDate: time.Date(2024, time.January, 22, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2024, time.January, 28,  0, 0, 0, 0, time.UTC),
		Capacity: 10,
	}

func TestGetBookings(t *testing.T) {
	bookingsService := NewBookingsService()

	bookRepoMock := new(repositories.MockBookingsRepository)
	bookRepoMock.On("GetBookings").Return(&[]dtos.BookingCompleteDTO{
		*firstbooking.ToBookingDTO(),
		*secondbooking.ToBookingDTO(),
	})

	bookingsService.BookingsRepository = bookRepoMock

	resultWithValues := bookingsService.GetBookings()

	require.Equal(t, []dtos.BookingCompleteDTO{
		*firstbooking.ToBookingDTO(),
		*secondbooking.ToBookingDTO(),
	}, *resultWithValues)
}

func TestGetBooking(t *testing.T){
	bookingsService := NewBookingsService()
	id := 1

	bookRepoMock := new(repositories.MockBookingsRepository)
	bookRepoMock.On("GetBooking", id).Return(&firstbooking, nil)

	bookingsService.BookingsRepository = bookRepoMock

	resultWithValues, errWithVals := bookingsService.GetBooking(id)

	require.Nil(t, errWithVals)
	require.Equal(t, firstbooking.ToBookingDTO(), resultWithValues)

	id = 0

	bookRepoMock.On("GetBooking", id).Return(nil, errors.New("error"))

	resultWithoutValues, errWithoutVals := bookingsService.GetBooking(id)

	require.Nil(t, resultWithoutValues)
	require.NotNil(t, errWithoutVals)
}

func TestInsertNewBooking(t *testing.T){
	bookingsService := NewBookingsService()

	repoMockInsertedBooking := &entities.Booking{
		BookingId: 3,
		Name:      "Jonas",
		Date:      time.Date(2024, time.January, 25, 0, 0, 0, 0, time.UTC),
		ClassId:   1,
	}

	 bookRepoMock := new(repositories.MockBookingsRepository)
	 bookRepoMock.On("InsertNewBooking", mock.AnythingOfType("*entities.Booking")).Return(repoMockInsertedBooking.ToBookingDTO())

	 classesRepoMock := new(classesRepo.MockClassesRepository)
	 classesRepoMock.On("GetClassSchedule", mock.AnythingOfType("int")).Return(&class, nil)

	 bookingsService.BookingsRepository = bookRepoMock
	 bookingsService.ClassesRepository = classesRepoMock

	 insertedBooking, errInsert := bookingsService.InsertNewBooking(&dtos.BookingDTO{
	 	Name: "Jonas",
	 	Date: time.Date(2024, time.January, 25, 0, 0, 0, 0, time.UTC),
	 	ClassId: 1,
	 })

	 require.Nil(t, errInsert)
	 require.Equal(t, dtos.BookingCompleteDTO{
	 	BookingId: 3,
	 	Name:      "Jonas",
	 	Date:      time.Date(2024, time.January, 25, 0, 0, 0, 0, time.UTC),
	 	ClassId:   1,
	 }, *insertedBooking)


	invalidBooking, err := bookingsService.InsertNewBooking(&dtos.BookingDTO{
		Name: "Jonas",
		Date: time.Date(2024, time.December, 25, 0, 0, 0, 0, time.UTC),
		ClassId: 1,
	})

	require.Nil(t, invalidBooking)
	require.NotNil(t, err)

	classesRepoMock = new(classesRepo.MockClassesRepository)
	classesRepoMock.On("GetClassSchedule", mock.AnythingOfType("int")).Return(nil, errors.New("error"))
	bookingsService.ClassesRepository = classesRepoMock

	bookingNoClass, errNoclass := bookingsService.InsertNewBooking(&dtos.BookingDTO{
		Name: "Jonas",
		Date: time.Date(2024, time.December, 25, 0, 0, 0, 0, time.UTC),
		ClassId: 1,
	})

	require.Nil(t, bookingNoClass)
	require.NotNil(t, errNoclass)
}

func TestUpdateBooking(t *testing.T) {
	bookingsService := NewBookingsService()

 	mockUpdatedBooking := entities.Booking{
	 	BookingId: 1,
	 	Name:      "Sam",
	 	Date:      time.Date(2024, time.January, 26, 0, 0, 0, 0, time.UTC),
	 	ClassId:   1,
	}

	bookRepoMock := new(repositories.MockBookingsRepository)
	bookRepoMock.On("GetBooking", mock.AnythingOfType("int")).Return(nil, errors.New("error"))
	bookingsService.BookingsRepository = bookRepoMock

	classesRepoMock := new(classesRepo.MockClassesRepository)
	classesRepoMock.On("GetClassSchedule", class.ClassId).Return(&class, nil)

	updatedBookingNoneExisting, errGet, errUpdate := bookingsService.UpdateBooking(1, &dtos.BookingDTO{
		Name:      "Sam",
	 	Date:      time.Date(2024, time.January, 26, 0, 0, 0, 0, time.UTC),
	 	ClassId:   1,
	})

	require.NotNil(t, errGet)
	require.Nil(t, errUpdate)
	require.Nil(t, updatedBookingNoneExisting)

	bookRepoMock = new(repositories.MockBookingsRepository)
	bookRepoMock.On("GetBooking", mock.AnythingOfType("int")).Return(&firstbooking, nil)
 	bookRepoMock.On("UpdateBooking", mock.AnythingOfType("int"), mock.AnythingOfType("*entities.Booking")).Return(mockUpdatedBooking.ToBookingDTO(), nil)

	bookingsService.BookingsRepository = bookRepoMock
	bookingsService.ClassesRepository = classesRepoMock

	updatedBooking, errGet, errUpdate := bookingsService.UpdateBooking(1, &dtos.BookingDTO{
		Name:      "Sam",
	 	Date:      time.Date(2024, time.January, 26, 0, 0, 0, 0, time.UTC),
	 	ClassId:   1,
	})

	require.Nil(t, errGet)
	require.Nil(t, errUpdate)
	require.Equal(t, mockUpdatedBooking.ToBookingDTO(), updatedBooking)

	updatedBookingInvalid, errGet, errUpdate := bookingsService.UpdateBooking(1, &dtos.BookingDTO{
		Name:      "Sam",
	 	Date:      time.Date(2024, time.December, 26, 0, 0, 0, 0, time.UTC),
	 	ClassId:   1,
	})

	require.Nil(t, errGet)
	require.NotNil(t, errUpdate)
	require.Nil(t, updatedBookingInvalid)
}

func TestDeleteBooking(t *testing.T){
	bookingsService := NewBookingsService()

	bookRepoMock := new(repositories.MockBookingsRepository)
	bookRepoMock.On("GetBooking", mock.AnythingOfType("int")).Return(&firstbooking, nil)
	bookRepoMock.On("DeleteBooking", mock.AnythingOfType("int")).Return(firstbooking.ToBookingDTO(), nil)
	bookingsService.BookingsRepository = bookRepoMock

	deletedBooking, errGet, errDelete := bookingsService.DeleteBooking(1)

	require.Nil(t, errGet)
	require.Nil(t, errDelete)
	require.Equal(t, firstbooking.ToBookingDTO(), deletedBooking)

	bookRepoMock = new(repositories.MockBookingsRepository)
	bookRepoMock.On("GetBooking", mock.AnythingOfType("int")).Return(nil, errors.New("error"))
	bookingsService.BookingsRepository = bookRepoMock

	deletedBooking, errGet, errDelete = bookingsService.DeleteBooking(0)

	require.NotNil(t, errGet)
	require.Nil(t, errDelete)
	require.Nil(t, deletedBooking)

	bookRepoMock = new(repositories.MockBookingsRepository)
	bookRepoMock.On("GetBooking", mock.AnythingOfType("int")).Return(&firstbooking, nil)
	bookRepoMock.On("DeleteBooking", mock.AnythingOfType("int")).Return(nil, errors.New("error"))
	bookingsService.BookingsRepository = bookRepoMock

	deletedBooking, errGet, errDelete = bookingsService.DeleteBooking(1)

	fmt.Println(errGet)

	require.Nil(t, errGet)
	require.Nil(t, deletedBooking)
	require.NotNil(t, errDelete)
}

func TestValidateBooking(t *testing.T){
	bookingsService := NewBookingsService()

	classesRepoMock := new(classesRepo.MockClassesRepository)
	classesRepoMock.On("GetClassSchedule", class.ClassId).Return(&class, nil)

	bookingsService.ClassesRepository = classesRepoMock

	err := bookingsService.validateBooking(&dtos.BookingDTO{
		Name:      "Sam",
	 	Date:      time.Date(2024, time.January, 26, 0, 0, 0, 0, time.UTC),
	 	ClassId:   1,
	})

	require.Nil(t, err)

	classesRepoMock = new(classesRepo.MockClassesRepository)
	classesRepoMock.On("GetClassSchedule", class.ClassId).Return(nil, errors.New("error"))
	bookingsService.ClassesRepository = classesRepoMock

	err = bookingsService.validateBooking(&dtos.BookingDTO{
		Name:      "Sam",
	 	Date:      time.Date(2024, time.January, 26, 0, 0, 0, 0, time.UTC),
	 	ClassId:   1,
	})

	require.NotNil(t, err)

	classesRepoMock = new(classesRepo.MockClassesRepository)
	classesRepoMock.On("GetClassSchedule", class.ClassId).Return(&class, nil)
	bookingsService.ClassesRepository = classesRepoMock

	err = bookingsService.validateBooking(&dtos.BookingDTO{
		Name:      "Sam",
	 	Date:      time.Date(2024, time.December, 26, 0, 0, 0, 0, time.UTC),
	 	ClassId:   1,
	})

	require.NotNil(t, err)
}