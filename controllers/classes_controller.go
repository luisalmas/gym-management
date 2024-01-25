package controllers

import (
	"gym-management/models/dtos"
	"gym-management/services"
	"net/http"
	"strconv"

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
	router.GET("/classes/:id", ctrl.getClassSchedule)
	router.POST("/classes", ctrl.postClassSchedule)
	router.PUT("/classes/:id", ctrl.putClassSchedule)
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

// GetClasses             godoc
// @Summary      Get class
// @Description  Returns single class.
// @Tags         classes
// @Produce      json
//@Param         id  path      string true  "ClassSchedule Id"
// @Success      200  {object}  entities.ClassSchedule
// @Error      	 404
// @Router       /classes/{id} [get]
func (ctrl *ClassesController) getClassSchedule(c *gin.Context) {
	id, idError := strconv.Atoi(c.Param("id"))
	if idError != nil {
		c.IndentedJSON(http.StatusBadRequest, idError.Error())
		return
	}
	class, err := ctrl.service.GetClassSchedule(id)

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
// @Success      201 {object} entities.ClassSchedule
// @Error      	 400 
// @Router       /classes [post]
func (ctrl *ClassesController) postClassSchedule(c *gin.Context) {
	var classSchedule dtos.ClassScheduleDTO

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

// GetClasses             godoc
// @Summary      Put classes
// @Description  Updates a single class.
// @Tags         classes
// @Produce      json
//@Param         id  path      string true  "ClassSchedule Id"
//@Param         classScheduleDTO  body      dtos.ClassScheduleDTO true  "ClassScheduleDTO JSON"
// @Success      200  {object}  entities.ClassSchedule
// @Error      	 400
// @Router       /classes/{id} [put]
func (ctrl *ClassesController) putClassSchedule(c *gin.Context) {
	
	classSchedule, bodyClassError := getClassScheduleFromRequest(c)
	if bodyClassError != nil {
        c.IndentedJSON(http.StatusBadRequest, bodyClassError.Error())
		return
    }

	id, idError := strconv.Atoi(c.Param("id"))
	if idError != nil {
		c.IndentedJSON(http.StatusBadRequest, idError.Error())
		return
	}

	//TODO fix this error when updating
	updatedClass, updateClassError := ctrl.service.UpdateClassSchedule(id, classSchedule)

	if updateClassError == nil{
		c.IndentedJSON(http.StatusNotFound, updateClassError.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, updatedClass)
}


func getClassScheduleFromRequest(c *gin.Context) (*dtos.ClassScheduleDTO, error){
	var classSchedule dtos.ClassScheduleDTO

	if err := c.BindJSON(&classSchedule); err != nil {
        c.IndentedJSON(http.StatusBadRequest, err.Error())
		return nil, err
    }

	return &classSchedule, nil
}