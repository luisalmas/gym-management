package controllers

import (
	"encoding/json"
	"gym-management/classes/models/dtos"
	"gym-management/classes/services"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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

	classesController := NewClassesController()
	router := gin.Default()
	prefix := router.Group("/api")
	
	classesController.SetupRoutes(prefix)

	classesServiceMock := new(services.MockClassesService)
	classesController.ClassesService = classesServiceMock

	//===================== GetClasses tests ==============================================

	t.Run("GetClasses", func(t *testing.T) {
		classesServiceMock.On("GetClasses").Return([]dtos.ClassCompleteDTO{*firstClass, *secondClass}).Once()

		request, _ := http.NewRequest("GET","/api/classes",nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		var responseData []dtos.ClassCompleteDTO
		err := json.Unmarshal(response.Body.Bytes(), &responseData)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Nil(t, err)
		assert.Equal(t, []dtos.ClassCompleteDTO{*firstClass, *secondClass}, responseData)
	})
}