package controller

import "github.com/gin-gonic/gin"

type newsController struct {
}

func NewNewsController() *newsController {
	return &newsController{}
}

func (nc *newsController) GetNews(c *gin.Context) {

}
