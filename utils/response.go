package utils

import (
	"github.com/gin-gonic/gin"
)

func Respond(c *gin.Context, status int, message string, result interface{}, err error) {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}

	c.JSON(status, gin.H{
		"message": message,
		"result":  result,
		"error":   errMsg,
	})
}
