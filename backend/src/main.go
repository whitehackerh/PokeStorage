package main

import (
	"github.com/gin-gonic/gin"
	"github.com/whitehackerh/PokeStorage/src/handler"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	api := r.Group("/api")
	{
		api.POST("/signup", handler.Signup)
	}
	r.Run()
}
