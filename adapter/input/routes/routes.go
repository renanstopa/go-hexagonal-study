package routes

import (
	"go-study/adapter/input/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	newsController := controller.NewNewsController()

	r.GET("/news", newsController.GetNews)
}
