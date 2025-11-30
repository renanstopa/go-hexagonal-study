package controller

import (
	"go-study/adapter/input/model/request"
	"go-study/application/domain"
	"go-study/application/port/input"
	"go-study/configuration/logger"
	"go-study/configuration/validation"
	"net/http"

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

	newsDomain := domain.NewsReqDomain{
		Subject: newsRequest.Subject,
		From:    newsRequest.From.Format("2006-01-02"),
	}

	newsResponseDomain, err := nc.newsUseCase.GetNewsService(newsDomain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, newsResponseDomain)
}
