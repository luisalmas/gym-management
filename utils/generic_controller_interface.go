package utils

import "github.com/gin-gonic/gin"

type GenericControllerInterface interface {
	SetupRoutes(router *gin.RouterGroup)
}