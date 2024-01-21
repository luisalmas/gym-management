package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClassesService struct {
}

func NewClassesService() *ClassesService {
	return &ClassesService{}
}

func (serv *ClassesService) GetItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "test")
}