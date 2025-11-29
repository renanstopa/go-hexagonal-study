package controller

import (
	"go-study/adapter/input/model/request"
	"go-study/application/port/input"
	"go-study/configuration/logger"
	"go-study/configuration/validation"

	"github.com/gin-gonic/gin"
)

type newsController struct {
	newsUseCase input.NewsUseCase
}

func NewNewsController(
	newsUseCase input.NewsUseCase,
) *newsController {
	return &newsController{newsUseCase}
}

func (nc *newsController) GetNews(c *gin.Context) {
	logger.Info("Inicializando GetNews controller API")
	newsRequest := request.NewsRequest{}

	if err := c.ShouldBindQuery(&newsRequest); err != nil {
		logger.Error("Erro ao tentar validar campos da requisição", err)
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	_, _ = nc.newsUseCase.GetNewsService(newsRequest.Subject, newsRequest.From)
}
