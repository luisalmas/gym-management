package controllers

import (
	"gym-management/classes/models/dtos"
	"gym-management/classes/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ClassesController struct {
	ClassesService *services.ClassesService
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
}

// GetClasse             godoc
// @Summary      Get classes
// @Description  Returns all scheduled classes.
// @Tags         classes
// @Produce      json
// @Success      200  {array}  dtos.ClassScheduleCompleteDTO
// @Router       /classes [get]
func (ctrl *ClassesController) getClassesSchedules(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, ctrl.ClassesService.GetClassesSchedules())
}

// GetClass             godoc
// @Summary      Get class
// @Description  Returns single class.
// @Tags         classes
// @Produce      json
//@Param         id  path      string true  "ClassSchedule Id"
// @Success      200  {object}  dtos.ClassScheduleCompleteDTO
// @Failure      404
// @Router       /classes/{id} [get]
func (ctrl *ClassesController) getClassSchedule(c *gin.Context) {
	id, idError := strconv.Atoi(c.Param("id"))
	if idError != nil {
		c.IndentedJSON(http.StatusBadRequest, idError.Error())
		return
	}
	class, err := ctrl.ClassesService.GetClassSchedule(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, class)
}

// PostClasses             godoc
// @Summary      Post class
// @Description  Post a new class.
// @Tags         classes
// @Produce      json
//@Param         classScheduleDTO  body      dtos.ClassScheduleDTO true  "ClassScheduleDTO JSON"
// @Success      201 {object} dtos.ClassScheduleCompleteDTO
// @Failure      400 
// @Router       /classes [post]
func (ctrl *ClassesController) postClassSchedule(c *gin.Context) {
	var classSchedule dtos.ClassScheduleDTO

	if err := c.BindJSON(&classSchedule); err != nil {
        c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
    }
	
	insertedClassSchedule, err := ctrl.ClassesService.InsertNewClassSchedule(&classSchedule)

	if err != nil{
		c.IndentedJSON(http.StatusBadRequest, err.Error())
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
//@Param         classScheduleDTO  body      dtos.ClassScheduleDTO true  "ClassScheduleDTO JSON"
// @Success      200  {object}  dtos.ClassScheduleCompleteDTO
// @Failure      400
// @Failure      404
// @Router       /classes/{id} [put]
func (ctrl *ClassesController) putClassSchedule(c *gin.Context) {
	
	var classSchedule dtos.ClassScheduleDTO

	if err := c.BindJSON(&classSchedule); err != nil {
        c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
    }

	id, idError := strconv.Atoi(c.Param("id"))
	if idError != nil {
		c.IndentedJSON(http.StatusBadRequest, idError.Error())
		return
	}

	updatedClass, unfoundError, updateError := ctrl.ClassesService.UpdateClassSchedule(id, &classSchedule)

	if unfoundError != nil{
		c.IndentedJSON(http.StatusNotFound, unfoundError.Error())
		return
	}

	if updateError != nil{
		c.IndentedJSON(http.StatusBadRequest, updateError.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, updatedClass)
}