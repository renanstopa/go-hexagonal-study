package input

import (
	"go-study/application/domain"
	"go-study/configuration/rest_err"
)

type NewsUseCase interface {
	GetNewsService(
		subject string, from string,
	) (*domain.NewsDomain, *rest_err.RestErr)
}
