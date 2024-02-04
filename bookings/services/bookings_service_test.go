package services

import (
	"gym-management/bookings/models/dtos"
	"gym-management/bookings/models/entities"
	"gym-management/bookings/models/errors"
	"gym-management/bookings/repositories"
	classesErrors "gym-management/classes/models/errors"

	classesEntities "gym-management/classes/models/entities"
	classesRepo "gym-management/classes/repositories"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBookingsService(t *testing.T) {
	var firstbooking = &entities.Booking{
		BookingId: 1,
		Name:      "Peter",
		Date:      time.Date(2024, time.January, 25, 0, 0, 0, 0, time.UTC),
		ClassId:   1,
	}
	
	var secondbooking = &entities.Booking{
		BookingId: 2,
		Name:      "Samantha",
		Date:      time.Date(2024, time.January, 26, 0, 0, 0, 0, time.UTC),
		ClassId:   1,
	}
	
	var bookingToInsert = &dtos.BookingDTO{
		Name: "Jonas",
		Date: time.Date(2024, time.January, 25, 0, 0, 0, 0, time.UTC),
		ClassId: 1,
	}

	var bookingAfterInsert = &dtos.BookingCompleteDTO{
		BookingId: 3,
		Name:      "Jonas",
		Date:      time.Date(2024, time.January, 25, 0, 0, 0, 0, time.UTC),
		ClassId:   1,
	}

	var bookingToInsertInvalidDate = &dtos.BookingDTO{
		Name: "Jonas",
		Date: time.Date(2024, time.December, 25, 0, 0, 0, 0, time.UTC),
		ClassId: 1,
	}

	var class = &classesEntities.Class{
			ClassId: 1,
			Name: "Class1",
			StartDate: time.Date(2024, time.January, 22, 0, 0, 0, 0, time.UTC),
			EndDate: time.Date(2024, time.January, 28,  0, 0, 0, 0, time.UTC),
			Capacity: 10,
		}

	bookingsService := NewBookingsService()
	bookRepoMock := new(repositories.MockBookingsRepository)
	bookingsService.BookingsRepository = bookRepoMock

	classesRepoMock := new(classesRepo.MockClassesRepository)
	bookingsService.ClassesRepository = classesRepoMock

	//===================== GetBookings tests ==============================================

	t.Run("GetBookings", func(t *testing.T) {
		bookRepoMock.On("GetBookings").Return(&[]dtos.BookingCompleteDTO{
			*firstbooking.ToBookingCompleteDTO(),
			*secondbooking.ToBookingCompleteDTO(),
		}).Once()

		resultWithValues := bookingsService.GetBookings()

		assert.Equal(t, []dtos.BookingCompleteDTO{
			*firstbooking.ToBookingCompleteDTO(),
			*secondbooking.ToBookingCompleteDTO(),
		}, *resultWithValues)
	})

	//===================== GetBooking tests ==============================================

	t.Run("GetBooking", func(t *testing.T){
		bookRepoMock.On("GetBooking", mock.Anything).Return(firstbooking, nil).Once()

		result, errWithVals := bookingsService.GetBooking(firstbooking.BookingId)

		assert.Nil(t, errWithVals)
		assert.Equal(t, firstbooking.ToBookingCompleteDTO(), result)
	})

	t.Run("GetBookingNotFound", func(t *testing.T){
		bookRepoMock.On("GetBooking", mock.Anything).Return(nil, errors.NewBookingNotFoundError()).Once()

		result, errWithVals := bookingsService.GetBooking(firstbooking.BookingId)

		assert.NotNil(t, errWithVals)
		assert.Nil(t, result)
	})

	//===================== InsertBooking tests ==============================================

	t.Run("InsertNewBooking", func(t *testing.T){
		bookRepoMock.On("InsertNewBooking", mock.Anything).Return(bookingAfterInsert).Once()
	 	classesRepoMock.On("GetClass", mock.AnythingOfType("int")).Return(class, nil).Once()

		insertedBooking, errInsert := bookingsService.InsertNewBooking(bookingToInsert)
   
		assert.Nil(t, errInsert)
		assert.Equal(t, *bookingAfterInsert, *insertedBooking)
	})

	t.Run("InsertBookingInvalidClass", func(t *testing.T){
		classesRepoMock.On("GetClass", mock.AnythingOfType("int")).Return(nil, classesErrors.NewClassNotFoundError()).Once()

		insertedBooking, errInsert := bookingsService.InsertNewBooking(bookingToInsert)
   
		assert.Nil(t, insertedBooking)
		assert.NotNil(t, errInsert)
	})

	t.Run("InsertBookingInvalidDate", func(t *testing.T){
		classesRepoMock.On("GetClass", mock.AnythingOfType("int")).Return(class, nil).Once()

		insertedBooking, errInsert := bookingsService.InsertNewBooking(bookingToInsertInvalidDate)

		assert.Nil(t, insertedBooking)
		assert.NotNil(t, errInsert)
	})

	//===================== UpdateBookings tests ==============================================

	t.Run("UpdateBooking", func(t *testing.T){
		bookRepoMock.On("GetBooking", mock.Anything).Return(firstbooking, nil).Once()
		classesRepoMock.On("GetClass", mock.AnythingOfType("int")).Return(class, nil).Once()
		bookRepoMock.On("UpdateBooking", mock.Anything, mock.Anything).Return(bookingAfterInsert).Once()

		updatedBooking, err := bookingsService.UpdateBooking(1, bookingToInsert)

		assert.Nil(t, err)
		assert.Equal(t, bookingAfterInsert, updatedBooking)
	})

	t.Run("UpdateBookingBookingNotFound", func(t *testing.T){
		bookRepoMock.On("GetBooking", mock.Anything).Return(nil, errors.NewBookingNotFoundError()).Once()
	
		 updatedBooking, err := bookingsService.UpdateBooking(1, bookingToInsert)

		 assert.NotNil(t, err)
		 assert.Nil(t, updatedBooking)
	})

	t.Run("UpdateBookingClassNotFound", func(t *testing.T){
		bookRepoMock.On("GetBooking", mock.Anything).Return(firstbooking, nil).Once()
		classesRepoMock.On("GetClass", mock.AnythingOfType("int")).Return(nil, classesErrors.NewClassNotFoundError()).Once()

		 updatedBooking, err := bookingsService.UpdateBooking(1, bookingToInsert)

		 assert.NotNil(t, err)
		 assert.Nil(t, updatedBooking)
	})

	t.Run("UpdateBookingInvalidDate", func(t *testing.T){
		bookRepoMock.On("GetBooking", mock.Anything).Return(firstbooking, nil).Once()
		classesRepoMock.On("GetClass", mock.AnythingOfType("int")).Return(class, nil).Once()

		updatedBooking, err := bookingsService.UpdateBooking(1, bookingToInsertInvalidDate)

		assert.NotNil(t, err)
		assert.Nil(t, updatedBooking)
	})

	//===================== DeleteBookings tests ==============================================

	t.Run("DeleteBooking", func(t *testing.T){
		bookRepoMock.On("DeleteBooking", mock.Anything).Return(firstbooking.ToBookingCompleteDTO(), nil).Once()

		deletedBooking, err := bookingsService.DeleteBooking(1)

		assert.Nil(t, err)
		assert.Equal(t, firstbooking.ToBookingCompleteDTO(), deletedBooking)
	})

	t.Run("DeleteBookingNotFound", func(t *testing.T){
		bookRepoMock.On("DeleteBooking", mock.Anything).Return(nil, errors.NewBookingNotFoundError()).Once()
		
		deletedBooking, err := bookingsService.DeleteBooking(1)

		assert.NotNil(t, err)
		assert.Nil(t, deletedBooking)
	})

	//===================== validateBooking tests ==============================================

	t.Run("validateBooking", func(t *testing.T){
		classesRepoMock.On("GetClass", class.ClassId).Return(class, nil).Once()

		err := bookingsService.validateBooking(bookingToInsert)

		assert.Nil(t, err)
	})

	t.Run("validateBookingClassNotFound", func(t *testing.T){
		classesRepoMock.On("GetClass", class.ClassId).Return(nil, classesErrors.NewClassNotFoundError()).Once()

		err := bookingsService.validateBooking(bookingToInsert)

		assert.NotNil(t, err)
	})

	t.Run("validateBookingInvalidDate", func(t *testing.T){
		classesRepoMock.On("GetClass", class.ClassId).Return(class, nil).Once()

		err := bookingsService.validateBooking(bookingToInsertInvalidDate)

		assert.NotNil(t, err)
	})

	bookRepoMock.AssertExpectations(t)
	classesRepoMock.AssertExpectations(t)
}