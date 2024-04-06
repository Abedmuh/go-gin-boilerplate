package routes

import (
	"abedmuh/go-gin-boilerplate/internal/app/home"

	"github.com/gin-gonic/gin"
)

func HomeRoutes(route *gin.RouterGroup) {
	controller := home.NewHomeController()

  route.GET("/", controller.GetHome)
}