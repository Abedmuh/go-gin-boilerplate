package main

import (
	"abedmuh/go-gin-boilerplate/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {


	app := gin.Default()

	v1 := app.Group("v1")
	{
		routes.HomeRoutes(v1)
	}

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	app.Run(":8080")
}