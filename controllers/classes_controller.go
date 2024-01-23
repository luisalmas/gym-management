package controllers

import (
	"gym-management/models/dtos"
	"gym-management/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClassesController struct {
	service *services.ClassesService
}

func NewClassesController() *ClassesController {
	return &ClassesController{
		service: services.NewClassesService(),
	}
}

func (ctrl *ClassesController) SetupRoutes(router *gin.RouterGroup) {
	router.GET("/classes", ctrl.getClassesSchedules)
	router.POST("/classes", ctrl.postClassSchedule)
	//router.PUT("/classes", ctrl.service.GetItems)
}

// GetClasses             godoc
// @Summary      Get classes
// @Description  Returns all scheduled classes.
// @Tags         classes
// @Produce      json
// @Success      200  {array}  entities.ClassSchedule
// @Router       /classes [get]
func (ctrl *ClassesController) getClassesSchedules(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, ctrl.service.GetClassesSchedules())
}

// PostClasses             godoc
// @Summary      Post class
// @Description  Post a new class.
// @Tags         classes
// @Produce      json
//@Param         classScheduleDTO  body      dtos.ClassScheduleDTO true  "ClassScheduleDTO JSON"
// @Success      201 {object} entities.ClassSchedule
// @Error      	 400 
// @Router       /classes [post]
func (ctrl *ClassesController) postClassSchedule(c *gin.Context) {
	var classSchedule dtos.ClassScheduleDTO

	//TODO fix bind
	if err := c.BindJSON(&classSchedule); err != nil {
        c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
    }
	
	insertedClassSchedule, err := ctrl.service.InsertNewClassSchedule(&classSchedule)

	if err != nil{
	
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusCreated, insertedClassSchedule)
}