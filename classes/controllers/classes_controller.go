package controllers

import (
	"gym-management/classes/models/dtos"
	"gym-management/classes/services"
	"gym-management/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
)

var (
	g = galidator.New()
	customizer = g.Validator(dtos.ClassDTO{})
  )

type ClassesController struct {
	ClassesService services.ClassesService
}

func NewClassesController() *ClassesController {
	return &ClassesController{
		ClassesService: services.NewClassesService(),
	}
}

func (ctrl *ClassesController) SetupRoutes(router *gin.RouterGroup) {
	router.GET("/classes", ctrl.getClassesSchedules)
	router.GET("/classes/:id", ctrl.getClassSchedule)
	router.POST("/classes", ctrl.postClassSchedule)
	router.PUT("/classes/:id", ctrl.putClassSchedule)
	router.DELETE("/classes/:id", ctrl.deleteClassSchedule)
	router.GET("/classes/:id/bookings", ctrl.getClassBookings)
}

// GetClasse             godoc
// @Summary      Get classes
// @Description  Returns all scheduled classes.
// @Tags         classes
// @Produce      json
// @Success      200  {array}  dtos.ClassCompleteDTO
// @Router       /classes [get]
func (ctrl *ClassesController) getClassesSchedules(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, ctrl.ClassesService.GetClasses())
}

// GetClass             godoc
// @Summary      Get class
// @Description  Returns single class.
// @Tags         classes
// @Produce      json
//@Param         id  path      string true  "ClassSchedule Id"
// @Success      200  {object}  dtos.ClassCompleteDTO
// @Failure      404
// @Router       /classes/{id} [get]
func (ctrl *ClassesController) getClassSchedule(c *gin.Context) {
	id, idError := strconv.Atoi(c.Param("id"))
	if idError != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": idError.Error()})
		return
	}
	class, err := ctrl.ClassesService.GetClass(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, class)
}

// PostClasses             godoc
// @Summary      Post class
// @Description  Post a new class.
// @Tags         classes
// @Produce      json
//@Param         ClassDTO  body      dtos.ClassDTO true  "ClassDTO JSON"
// @Success      201 {object} dtos.ClassCompleteDTO
// @Failure      400 
// @Router       /classes [post]
func (ctrl *ClassesController) postClassSchedule(c *gin.Context) {
	var classSchedule dtos.ClassDTO

	if err := c.BindJSON(&classSchedule); err != nil {

		if e, _ := err.(*time.ParseError); e != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid date format"})
			return
		}

        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": customizer.DecryptErrors(err)})
		return
    }
	
	insertedClassSchedule, err := ctrl.ClassesService.InsertNewClass(&classSchedule)

	if err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, insertedClassSchedule)
}

// PutClasses             godoc
// @Summary      Put classes
// @Description  Updates a class.
// @Tags         classes
// @Produce      json
//@Param         id  path      string true  "ClassSchedule Id"
//@Param         ClassDTO  body      dtos.ClassDTO true  "ClassDTO JSON"
// @Success      200  {object}  dtos.ClassCompleteDTO
// @Failure      400
// @Failure      404
// @Router       /classes/{id} [put]
func (ctrl *ClassesController) putClassSchedule(c *gin.Context) {
	
	var classSchedule dtos.ClassDTO

	if err := c.BindJSON(&classSchedule); err != nil {

		if e, _ := err.(*time.ParseError); e != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid date format"})
			return
		}

        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": customizer.DecryptErrors(err)})
		return
    }

	id, idError := strconv.Atoi(c.Param("id"))
	if idError != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": idError.Error()})
		return
	}

	updatedClass, updateError := ctrl.ClassesService.UpdateClass(id, &classSchedule)

	if updateError != nil{
		utils.ErrorHandler(c, updateError)
		return
	}

	c.IndentedJSON(http.StatusOK, updatedClass)
}

// DelteClasses             godoc
// @Summary      Delete classes
// @Description  Deletes a class.
// @Tags         classes
// @Produce      json
//@Param         id  path      string true  "ClassSchedule Id"
// @Success      200  {object}  dtos.ClassCompleteDTO
// @Failure      400
// @Failure      404
// @Router       /classes/{id} [delete]
func (ctrl *ClassesController) deleteClassSchedule(c *gin.Context) {
	id, idError := strconv.Atoi(c.Param("id"))
	if idError != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": idError.Error()})
		return
	}

	deletedClass, deleteError := ctrl.ClassesService.DeleteClass(id)

	if deleteError != nil{
		utils.ErrorHandler(c, deleteError)
		return
	}

	c.IndentedJSON(http.StatusOK, deletedClass)
}

// GetClass             godoc
// @Summary      Get class bookings
// @Description  Returns the bookings of a class.
// @Tags         classes
// @Produce      json
//@Param         id  path      string 	true  	"ClassSchedule Id"
//@Param    	date query 		string	false	"Class date in RFC3339"	
// @Success      200  {array}  dtos.BookingCompleteDTO
// @Failure      404
// @Router       /classes/{id}/bookings [get]
func (ctrl *ClassesController) getClassBookings(c *gin.Context){
	id, idError := strconv.Atoi(c.Param("id"))
	if idError != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": idError.Error()})
		return
	}

	dateStr := c.DefaultQuery("date", "")
	var date time.Time
	var err error

	
	if dateStr != "" {
		if date , err = time.Parse(time.RFC3339, dateStr) ; err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid date format"})
			return
		}
	}else{
		date = time.Time{}
	}

	bookings, err := ctrl.ClassesService.GetBookingsFromClass(id, date)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, bookings)
}	