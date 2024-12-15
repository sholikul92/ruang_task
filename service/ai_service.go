package service

import (
	airepository "a21hc3NpZ25tZW50/repository"
	"errors"
)

type AiServiceInterface interface {
	GeneratedText(modelPath, input string) (string, error)
	AnalyzeData(modelPath string, table map[string][]string, query string) (string, error)
}

type aiservice struct {
	Repo airepository.AiRepositoryInterface
}

func NewAiService(repo airepository.AiRepositoryInterface) *aiservice {
	return &aiservice{Repo: repo}
}

func (s *aiservice) GeneratedText(modelPath, input string) (string, error) {
	if modelPath == "" || input == "" {
		return "", errors.New("model path or input cannot empty")
	}

	textGenerate, err := s.Repo.TextGeneration(modelPath, input)
	if err != nil {
		return "", err
	}

	if len(textGenerate.Choices) == 0 {
		return "", errors.New("missing content")
	}

	var result = textGenerate.Choices[0].Messages.Content

	return result, nil
}

func (s *aiservice) AnalyzeData(modelPath string, table map[string][]string, query string) (string, error) {
	if len(table) == 0 {
		return "", errors.New("table is empty")
	}

	res, err := s.Repo.TableQuestionAnswering(modelPath, table, query)
	if err != nil {
		return "", err
	}

	if len(res.Cells) == 0 {
		return "", errors.New("cells is empty")
	}

	return res.Cells[0], nil
}
