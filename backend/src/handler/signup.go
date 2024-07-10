package handler

import (
	// usecase "../UseCases"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test",
	})
	// requestbody bind
	// validate
	// call usecase
}
