package services

import (
	bookingsDtos "gym-management/bookings/models/dtos"
	bookingsRepo "gym-management/bookings/repositories"
	"gym-management/classes/models/dtos"
	"gym-management/classes/models/entities"
	"gym-management/classes/models/errors"
	"gym-management/classes/repositories"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestClassesService(t *testing.T) {
	var firstClass = &entities.Class{
			ClassId: 1,
			Name: "Class1",
			StartDate: time.Date(2024, time.January, 22, 0, 0, 0, 0, time.UTC),
			EndDate: time.Date(2024, time.January, 28,  0, 0, 0, 0, time.UTC),
			Capacity: 10,
		}

	var secondClass = &entities.Class{
			ClassId: 2,
			Name: "Class2",
			StartDate: time.Date(2024, time.January, 29, 0, 0, 0, 0, time.UTC),
			EndDate: time.Date(2024, time.January, 31,  0, 0, 0, 0, time.UTC),
			Capacity: 20,
		}

	var insertClass = &entities.Class{
		ClassId: 0,
		Name: "Class3",
		StartDate: time.Date(2024, time.January, 01, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2024, time.January, 20,  0, 0, 0, 0, time.UTC),
		Capacity: 30,
	}

	var updateClass = &entities.Class{
		ClassId: 1,
		Name: "Class3",
		StartDate: time.Date(2024, time.January, 01, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2024, time.January, 20,  0, 0, 0, 0, time.UTC),
		Capacity: 30,
	}

	var firstClassBookings = &[]bookingsDtos.BookingCompleteDTO{
		{
			BookingId: 3,
			Name:      "Jonas",
			Date:      time.Date(2024, time.January, 25, 0, 0, 0, 0, time.UTC),
			ClassId:   1,
		},
		{
			BookingId: 4,
			Name:      "Joana",
			Date:      time.Date(2024, time.January, 26, 0, 0, 0, 0, time.UTC),
			ClassId:   1,
		},
	}

	classesService := NewClassesService()
	classesRepoMock := new(repositories.MockClassesRepository)
	classesService.ClassesRepository = classesRepoMock
	bookingsRepoMock := new(bookingsRepo.MockBookingsRepository)
	classesService.BookingsRepository = bookingsRepoMock

	//===================== GetClasses tests ==============================================

	t.Run("GetClasses", func(t *testing.T){
		classesRepoMock.On("GetClasses").Return(&[]dtos.ClassCompleteDTO{
			*firstClass.ToClassCompleteDTO(),
			*secondClass.ToClassCompleteDTO(),
		}).Once()

		classes := classesService.GetClasses()

		assert.Equal(t, &[]dtos.ClassCompleteDTO{
			*firstClass.ToClassCompleteDTO(),
			*secondClass.ToClassCompleteDTO(),
			}, classes)
	})

	//===================== GetClass tests ==============================================

	t.Run("GetClass", func(t *testing.T){
		classesRepoMock.On("GetClass", mock.Anything).Return(firstClass, nil).Once()

		class, err := classesService.GetClass(1)

		assert.Nil(t, err)
		assert.Equal(t, firstClass.ToClassCompleteDTO(), class)
	})

	t.Run("GetClassNotFound", func(t *testing.T){
		classesRepoMock.On("GetClass", mock.Anything).Return(nil, errors.NewClassNotFoundError()).Once()

		class, err := classesService.GetClass(1)

		assert.NotNil(t, err)
		assert.Nil(t, class)
	})

	//===================== InsertClass tests ==============================================

	t.Run("InsertClass", func(t *testing.T) {
		classesRepoMock.On("InsertNewClass", insertClass).Return(insertClass.ToClassCompleteDTO()).Once()

		class, err := classesService.InsertNewClass(&dtos.ClassDTO{
			Name: "Class3",
			StartDate: time.Date(2024, time.January, 01, 0, 0, 0, 0, time.UTC),
			EndDate: time.Date(2024, time.January, 20,  0, 0, 0, 0, time.UTC),
			Capacity: 30,
		})

		assert.Nil(t, err)
		assert.Equal(t, insertClass.ToClassCompleteDTO(), class)
	})

	t.Run("InsertClassInvalidCapacity", func(t *testing.T) {
		class, err := classesService.InsertNewClass(&dtos.ClassDTO{
			Name: "Class3",
			StartDate: time.Date(2024, time.February, 01, 0, 0, 0, 0, time.UTC),
			EndDate: time.Date(2024, time.January, 20,  0, 0, 0, 0, time.UTC),
			Capacity: -1,
		})

		assert.NotNil(t, err)
		assert.Nil(t, class)
	})

	t.Run("InsertClassInvalidDates", func(t *testing.T) {
		class, err := classesService.InsertNewClass(&dtos.ClassDTO{
			Name: "Class3",
			StartDate: time.Date(2024, time.February, 01, 0, 0, 0, 0, time.UTC),
			EndDate: time.Date(2023, time.January, 20,  0, 0, 0, 0, time.UTC),
			Capacity: -1,
		})

		assert.NotNil(t, err)
		assert.Nil(t, class)
	})

	t.Run("InsertClassInvalidName", func(t *testing.T) {
		class, err := classesService.InsertNewClass(&dtos.ClassDTO{
			Name: " ",
			StartDate: time.Date(2024, time.January, 01, 0, 0, 0, 0, time.UTC),
			EndDate: time.Date(2023, time.January, 20,  0, 0, 0, 0, time.UTC),
			Capacity: -1,
		})

		assert.NotNil(t, err)
		assert.Nil(t, class)
	})

	t.Run("InsertClassNameTrim", func(t *testing.T) {
		classesRepoMock.On("InsertNewClass", mock.Anything).Return(&dtos.ClassCompleteDTO{
			ClassId: 3,
			Name: "Name",
			StartDate: time.Date(2024, time.January, 01, 0, 0, 0, 0, time.UTC),
			EndDate: time.Date(2024, time.January, 20,  0, 0, 0, 0, time.UTC),
			Capacity: 1,
		}).Once()
		
		class, err := classesService.InsertNewClass(&dtos.ClassDTO{
			Name: "  Name  ",
			StartDate: time.Date(2024, time.January, 01, 0, 0, 0, 0, time.UTC),
			EndDate: time.Date(2024, time.January, 20,  0, 0, 0, 0, time.UTC),
			Capacity: 1,
		})

		assert.Nil(t, err)
		assert.Equal(t, dtos.ClassCompleteDTO{
			ClassId: 3,
			Name: "Name",
			StartDate: time.Date(2024, time.January, 01, 0, 0, 0, 0, time.UTC),
			EndDate: time.Date(2024, time.January, 20,  0, 0, 0, 0, time.UTC),
			Capacity: 1,
		}, *class)
	})

	//===================== UpdateClass tests ==============================================

	t.Run("UpdateClass", func(t *testing.T) {
		classesRepoMock.On("GetClass", updateClass.ClassId).Return(firstClass, nil).Once()
		classesRepoMock.On("UpdateClass", updateClass.ClassId, updateClass).Return(updateClass.ToClassCompleteDTO()).Once()
		bookingsRepoMock.On("DeleteBookingsFromClass", updateClass.ClassId, updateClass.StartDate, updateClass.EndDate).Return(firstClassBookings).Once()

		class, err := classesService.UpdateClass(updateClass.ClassId, &dtos.ClassDTO{
			Name: "Class3",
			StartDate: time.Date(2024, time.January, 01, 0, 0, 0, 0, time.UTC),
			EndDate: time.Date(2024, time.January, 20,  0, 0, 0, 0, time.UTC),
			Capacity: 30,
		})

		assert.Nil(t, err)
		assert.Equal(t, updateClass.ToClassCompleteDTO(), class)
	})

	t.Run("UpdateClassNotFound", func(t *testing.T) {
		classesRepoMock.On("GetClass", mock.Anything).Return(nil, errors.NewClassNotFoundError()).Once()

		class, err := classesService.UpdateClass(updateClass.ClassId, &dtos.ClassDTO{
			Name: "Class3",
			StartDate: time.Date(2024, time.February, 01, 0, 0, 0, 0, time.UTC),
			EndDate: time.Date(2024, time.January, 20,  0, 0, 0, 0, time.UTC),
			Capacity: 30,
		})

		assert.NotNil(t, err)
		assert.Nil(t, class)
	})

	t.Run("UpdateClassInvalidCapacity", func(t *testing.T) {
		classesRepoMock.On("GetClass", mock.Anything).Return(firstClass, nil).Once()

		class, err := classesService.UpdateClass(updateClass.ClassId, &dtos.ClassDTO{
			Name: "Class3",
			StartDate: time.Date(2024, time.February, 01, 0, 0, 0, 0, time.UTC),
			EndDate: time.Date(2024, time.January, 20,  0, 0, 0, 0, time.UTC),
			Capacity: -1,
		})

		assert.NotNil(t, err)
		assert.Nil(t, class)
	})

	t.Run("UpdateClassInvalidDates", func(t *testing.T) {
		classesRepoMock.On("GetClass", mock.Anything).Return(firstClass, nil).Once()
		
		class, err := classesService.UpdateClass(updateClass.ClassId, &dtos.ClassDTO{
			Name: "Class3",
			StartDate: time.Date(2024, time.February, 01, 0, 0, 0, 0, time.UTC),
			EndDate: time.Date(2023, time.January, 20,  0, 0, 0, 0, time.UTC),
			Capacity: 30,
		})

		assert.NotNil(t, err)
		assert.Nil(t, class)
	})

	//===================== DeleteClass tests ==============================================

	t.Run("DeleteClass", func(t *testing.T) {
		classesRepoMock.On("DeleteClass", firstClass.ClassId).Return(firstClass.ToClassCompleteDTO(), nil).Once()
		bookingsRepoMock.On("DeleteBookingsFromClass", firstClass.ClassId, time.Time{}, time.Time{}).Return(firstClassBookings).Once()


		class, err := classesService.DeleteClass(firstClass.ClassId)

		assert.Nil(t, err)
		assert.Equal(t, firstClass.ToClassCompleteDTO(), class)

	})

	t.Run("DeleteClassNotFound", func(t *testing.T) {
		classesRepoMock.On("DeleteClass", mock.Anything).Return(nil, errors.NewClassNotFoundError()).Once()
		
		class, err := classesService.DeleteClass(firstClass.ClassId)

		assert.NotNil(t, err)
		assert.Nil(t, class)
	})

	//===================== GetBookingsFromClass tests ==============================================

	t.Run("GetBookingsFromClass", func(t *testing.T){
		classesRepoMock.On("GetClass", firstClass.ClassId).Return(firstClass, nil).Once()
		bookingsRepoMock.On("GetBookingsFromClass", firstClass.ClassId, time.Time{}).Return(firstClassBookings).Once()

		bookings, err := classesService.GetBookingsFromClass(firstClass.ClassId, time.Time{})

		assert.Nil(t, err)
		assert.Equal(t, firstClassBookings, bookings)
	})

	t.Run("GetBookingsFromClassClassNotFound", func(t *testing.T){
		classesRepoMock.On("GetClass", mock.Anything).Return(nil, errors.NewClassNotFoundError()).Once()

		bookings, err := classesService.GetBookingsFromClass(1, time.Time{})

		assert.NotNil(t, err)
		assert.Nil(t, bookings)
	})

	classesRepoMock.AssertExpectations(t)
	bookingsRepoMock.AssertExpectations(t)
}