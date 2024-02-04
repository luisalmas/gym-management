package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Error interface {
	ErrorCode() int
	Error() string
}

func ErrorHandler(c *gin.Context, err error) {
	var errInterface Error
	if errors.As(err, &errInterface) {
		c.IndentedJSON(errInterface.ErrorCode(), gin.H{"message": errInterface.Error()})
		return
	}
	c.IndentedJSON(http.StatusBadRequest, err.Error())
}