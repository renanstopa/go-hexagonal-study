package routes

import (
	"go-study/adapter/input/controller"
	"go-study/application/service"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	newsService := service.NewNewsService()
	newsController := controller.NewNewsController(newsService)

	r.GET("/news", newsController.GetNews)
}
