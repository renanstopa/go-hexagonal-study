package service

import (
	"fmt"
	"go-study/application/domain"
	"go-study/application/port/output"
	"go-study/configuration/logger"
	"go-study/configuration/rest_err"
)

type newsService struct {
	newsPort output.NewsPort
}

func NewNewsService(
	newsPort output.NewsPort,
) *newsService {
	return &newsService{newsPort}
}

func (ns *newsService) GetNewsService(
	new domain.NewsReqDomain,
) (*domain.NewsDomain, *rest_err.RestErr) {
	logger.Info(fmt.Sprintf("Inicializando função getNewsService, subject=%s, from=%s", new.Subject, new.From))

	return ns.newsPort.GetNewsPort(new)
}
