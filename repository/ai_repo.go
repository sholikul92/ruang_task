package repository

import (
	"a21hc3NpZ25tZW50/model"

	"github.com/go-resty/resty/v2"
)

type AiRepositoryInterface interface {
	TextGeneration(modelPath, input string) (model.TextGenerationResponse, error)
	TableQuestionAnswering(modelPath string, table map[string][]string, query string) (model.TapasResponse, error)
}

type aiRepository struct {
	Client  *resty.Client
	ApiKey  string
	BaseUrl string
}

func NewAiRepository(client *resty.Client, apiKey, baseUrl string) *aiRepository {
	return &aiRepository{
		Client:  client,
		ApiKey:  apiKey,
		BaseUrl: baseUrl,
	}
}

func (repo *aiRepository) TextGeneration(modelPath, input string) (model.TextGenerationResponse, error) {
	var payload = model.TextGenerationPayload{
		Model: modelPath,
		Messages: []model.Messages{
			{
				Role:    "user",
				Content: input,
			},
		},
		MaxContents: 500,
		Stream:      false,
	}

	var response model.TextGenerationResponse

	_, err := repo.Client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+repo.ApiKey).
		SetBody(payload).
		SetResult(&response).
		Post(repo.BaseUrl + modelPath + "/v1/chat/completions")

	if err != nil {
		return model.TextGenerationResponse{}, err
	}

	return response, nil
}

func (repo *aiRepository) TableQuestionAnswering(modelPath string, table map[string][]string, query string) (model.TapasResponse, error) {
	var payload = model.AIRequest{
		Inputs: model.Inputs{
			Table: table,
			Query: query,
		},
	}

	var response model.TapasResponse

	_, err := repo.Client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+repo.ApiKey).
		SetBody(payload).
		SetResult(&response).
		Post(repo.BaseUrl + modelPath)

	if err != nil {
		return model.TapasResponse{}, err
	}

	return response, nil
}
