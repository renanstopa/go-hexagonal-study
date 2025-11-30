package routes

import (
	"go-study/adapter/input/controller"
	"go-study/adapter/output/news_http"
	"go-study/application/service"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	newsClient := news_http.NewNewsClient()
	newsService := service.NewNewsService(newsClient)
	newsController := controller.NewNewsController(newsService)

	r.GET("/news", newsController.GetNews)
}
