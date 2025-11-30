package news_http

import (
	"go-study/adapter/output/model/response"
	"go-study/application/domain"
	"go-study/configuration/env"
	"go-study/configuration/rest_err"

	"github.com/go-resty/resty/v2"
	"github.com/jinzhu/copier"
)

type newsClient struct{}

func NewNewsClient() *newsClient {
	client = resty.New().SetBaseURL("https://newsapi.org/v2")
	return &newsClient{}
}

var (
	client *resty.Client
)

func (nc *newsClient) GetNewsPort(
	newsDomain domain.NewsReqDomain) (*domain.NewsDomain, *rest_err.RestErr) {

	newsResponse := &response.NewsClientResponse{}

	_, err := client.R().
		SetQueryParams(map[string]string{
			"q":      newsDomain.Subject,
			"from":   newsDomain.From,
			"apiKey": env.GetNewsTokenApi(),
		}).
		SetResult(newsResponse).
		Get("/everything")

	if err != nil {
		return nil, rest_err.NewInternalServerError("Error trying to call NewsAPI with params")
	}

	newsResponseDomain := &domain.NewsDomain{}
	copier.Copy(newsResponseDomain, newsResponse)

	return newsResponseDomain, nil
}
