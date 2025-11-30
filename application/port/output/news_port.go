package output

import (
	"go-study/application/domain"
	"go-study/configuration/rest_err"
)

type NewsPort interface {
	GetNewsPort(
		domain.NewsReqDomain,
	) (*domain.NewsDomain, *rest_err.RestErr)
}
