package service

import (
	"fmt"
	"go-study/application/domain"
	"go-study/configuration/logger"
	"go-study/configuration/rest_err"
)

type newsService struct{}

func NewNewsService() *newsService {
	return &newsService{}
}

func (*newsService) GetNewsService(
	subject string, from string,
) (*domain.NewsDomain, *rest_err.RestErr) {
	logger.Info(fmt.Sprintf("Inicializando função getNewsService, subject=%s, from=%s", subject, from))

	return nil, nil
}
