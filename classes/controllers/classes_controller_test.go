package controllers

import (
	"encoding/json"
	bookingsDtos "gym-management/bookings/models/dtos"
	bookingsErrors "gym-management/bookings/models/errors"
	"gym-management/classes/models/dtos"
	"gym-management/classes/models/errors"
	"gym-management/classes/services"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestClassesController(t *testing.T) {

	var firstClass = &dtos.ClassCompleteDTO{
		ClassId: 1,
		Name: "Class1",
		StartDate: time.Date(2024, time.January, 22, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2024, time.January, 28,  0, 0, 0, 0, time.UTC),
		Capacity: 10,
	}

	var secondClass = &dtos.ClassCompleteDTO{
		ClassId: 2,
		Name: "Class2",
		StartDate: time.Date(2024, time.January, 29, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2024, time.January, 31,  0, 0, 0, 0, time.UTC),
		Capacity: 20,
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

	var insertClass = &dtos.ClassDTO{
		Name: "Class3",
		StartDate: time.Date(2024, time.January, 01, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2024, time.January, 20,  0, 0, 0, 0, time.UTC),
		Capacity: 30,
	}

	var insertedClass = &dtos.ClassCompleteDTO{
		ClassId: 3,
		Name: "Class3",
		StartDate: time.Date(2024, time.January, 01, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2024, time.January, 20,  0, 0, 0, 0, time.UTC),
		Capacity: 30,
	}

	classesController := NewClassesController()
	router := gin.Default()
	prefix := router.Group("/api")
	
	classesController.SetupRoutes(prefix)

	classesServiceMock := new(services.MockClassesService)
	classesController.ClassesService = classesServiceMock

	//===================== GetClasses tests ==============================================

	t.Run("GetClasses", func(t *testing.T) {
		classesServiceMock.On("GetClasses").Return(&[]dtos.ClassCompleteDTO{*firstClass, *secondClass}).Once()

		request, _ := http.NewRequest("GET","/api/classes",nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		var responseData []dtos.ClassCompleteDTO 
		err := json.Unmarshal(response.Body.Bytes(), &responseData)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Nil(t, err)
		assert.Equal(t, []dtos.ClassCompleteDTO{*firstClass, *secondClass}, responseData)
	})

	//===================== GetClass tests ==============================================

	t.Run("GetClass", func(t *testing.T) {
		classesServiceMock.On("GetClass", mock.Anything).Return(firstClass, nil).Once()

		request, _ := http.NewRequest("GET","/api/classes/1",nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		var responseData dtos.ClassCompleteDTO 
		err := json.Unmarshal(response.Body.Bytes(), &responseData)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Nil(t, err)
		assert.Equal(t, *firstClass, responseData)
	})

	t.Run("GetClassInvalidId", func(t *testing.T) {
		request, _ := http.NewRequest("GET","/api/classes/asd",nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("GetClassNotFound", func(t *testing.T) {
		classesServiceMock.On("GetClass", mock.Anything).Return(nil, errors.NewClassNotFoundError()).Once()

		request, _ := http.NewRequest("GET","/api/classes/1",nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	//===================== PostClass tests ==============================================

	t.Run("PostClass", func(t *testing.T) {
		classesServiceMock.On("InsertNewClass", mock.Anything).Return(insertedClass, nil).Once()

		jsonData, errorJson := json.Marshal(*insertClass)

		assert.Nil(t, errorJson)

		request, _ := http.NewRequest("POST","/api/classes", strings.NewReader(string(jsonData)))
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		var responseData dtos.ClassCompleteDTO 
		err := json.Unmarshal(response.Body.Bytes(), &responseData)

		assert.Equal(t, http.StatusCreated, response.Code)
		assert.Nil(t, err)
		assert.Equal(t, *insertedClass, responseData)
	})

	t.Run("PostClassNoBody", func(t *testing.T){
		request, _ := http.NewRequest("POST","/api/classes", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("PostClassInsertError", func(t *testing.T){
		classesServiceMock.On("InsertNewClass", mock.Anything).Return(nil, errors.NewClassInvalidError("error")).Once()

		jsonData, errorJson := json.Marshal(*insertClass)

		assert.Nil(t, errorJson)

		request, _ := http.NewRequest("POST","/api/classes", strings.NewReader(string(jsonData)))
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	//===================== PutClass tests ==============================================

	t.Run("PutClass", func(t *testing.T) {
		classesServiceMock.On("UpdateClass", mock.Anything, mock.Anything).Return(&dtos.ClassCompleteDTO{
			ClassId: 1,
			Name: "Class3",
			StartDate: time.Date(2024, time.January, 01, 0, 0, 0, 0, time.UTC),
			EndDate: time.Date(2024, time.January, 20,  0, 0, 0, 0, time.UTC),
			Capacity: 30,
			}, nil).Once()

		jsonData, errorJson := json.Marshal(*insertClass)

		assert.Nil(t, errorJson)

		request, _ := http.NewRequest("PUT","/api/classes/1", strings.NewReader(string(jsonData)))
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		var responseData dtos.ClassCompleteDTO 
		err := json.Unmarshal(response.Body.Bytes(), &responseData)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, dtos.ClassCompleteDTO{
			ClassId: 1,
			Name: "Class3",
			StartDate: time.Date(2024, time.January, 01, 0, 0, 0, 0, time.UTC),
			EndDate: time.Date(2024, time.January, 20,  0, 0, 0, 0, time.UTC),
			Capacity: 30,
			}, responseData)
	})

	t.Run("PutClassInvalidId", func(t *testing.T) {

		jsonData, errorJson := json.Marshal(*insertClass)

		assert.Nil(t, errorJson)

		request, _ := http.NewRequest("PUT","/api/classes/asd", strings.NewReader(string(jsonData)))
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("PutClassNoBody", func(t *testing.T) {
		request, _ := http.NewRequest("PUT","/api/classes/asd", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("PutClassNotFound", func(t *testing.T) {
		classesServiceMock.On("UpdateClass", mock.Anything, mock.Anything).Return(nil, errors.NewClassNotFoundError()).Once()

		jsonData, errorJson := json.Marshal(*insertClass)

		assert.Nil(t, errorJson)

		request, _ := http.NewRequest("PUT","/api/classes/1", strings.NewReader(string(jsonData)))
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("PutClassInvalid", func(t *testing.T) {
		classesServiceMock.On("UpdateClass", mock.Anything, mock.Anything).Return(nil, errors.NewClassInvalidError("error")).Once()

		jsonData, errorJson := json.Marshal(*insertClass)

		assert.Nil(t, errorJson)

		request, _ := http.NewRequest("PUT","/api/classes/1", strings.NewReader(string(jsonData)))
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	//===================== DeleteClass tests ==============================================

	t.Run("DeleteClass", func(t *testing.T){
		classesServiceMock.On("DeleteClass", mock.Anything).Return(firstClass, nil).Once()

		request, _ := http.NewRequest("DELETE","/api/classes/1", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		var responseData dtos.ClassCompleteDTO 
		err := json.Unmarshal(response.Body.Bytes(), &responseData)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, *firstClass, responseData)
	})

	t.Run("DeleteClasssInvalidId", func(t *testing.T) {
		request, _ := http.NewRequest("DELETE","/api/classes/asd", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("DeleteClassNotFound", func(t *testing.T) {
		classesServiceMock.On("DeleteClass", mock.Anything).Return(nil, errors.NewClassNotFoundError()).Once()

		request, _ := http.NewRequest("DELETE","/api/classes/1", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	//===================== GetBookingsOfClass tests ==============================================

	t.Run("GetBookingsFromClass", func(t *testing.T) {
		classesServiceMock.On("GetBookingsFromClass", mock.Anything, mock.Anything).Return(firstClassBookings, nil).Once()

		request, _ := http.NewRequest("GET","/api/classes/1/bookings", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		var responseData []bookingsDtos.BookingCompleteDTO 
		err := json.Unmarshal(response.Body.Bytes(), &responseData)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, *firstClassBookings, responseData)
	})

	t.Run("GetBookingsFromClassWithDate", func(t *testing.T) {
		classesServiceMock.On("GetBookingsFromClass", mock.Anything, mock.Anything).Return(firstClassBookings, nil).Once()

		request, _ := http.NewRequest("GET","/api/classes/1/bookings?date=2024-01-26T00:00:00Z", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		var responseData []bookingsDtos.BookingCompleteDTO 
		err := json.Unmarshal(response.Body.Bytes(), &responseData)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, *firstClassBookings, responseData)
	})

	t.Run("GetBookingsFromClassInvalidId", func(t *testing.T) {
		request, _ := http.NewRequest("GET","/api/classes/asd/bookings?date=2024-01-26T00:00:00Z", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("GetBookingsFromClassInvalidDate", func(t *testing.T) {
		request, _ := http.NewRequest("GET","/api/classes/1/bookings?date=2024", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("GetBookingsFromClassNotFound", func(t *testing.T) {
		classesServiceMock.On("GetBookingsFromClass", mock.Anything, mock.Anything).Return(nil, bookingsErrors.NewBookingNotFoundError()).Once()
		request, _ := http.NewRequest("GET","/api/classes/1/bookings", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)
		
		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	classesServiceMock.AssertExpectations(t)
}