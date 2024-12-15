package service

import (
	repository "a21hc3NpZ25tZW50/repository"
	"errors"
	"strings"
)

type FileServiceInterface interface {
	ProcessFile(fileContent string) (map[string][]string, error)
}

type fileService struct {
	Repo repository.FileRepository
}

func NewFileService(repo repository.FileRepository) *fileService {
	return &fileService{Repo: repo}
}

func (s *fileService) ProcessFile(fileContent string) (map[string][]string, error) {
	if len(fileContent) == 0 {
		return nil, errors.New("content is empty")
	}

	if len(fileContent) < 2 {
		return nil, errors.New("missing headers")
	}

	fileContentSplit := strings.Split(fileContent, "\n")
	headers := fileContentSplit[0]
	headerSplit := strings.Split(headers, ",")
	if len(headerSplit) == 0 {
		return nil, errors.New("invalid format header")
	}

	contents := []string{}
	for _, line := range fileContentSplit[1:] {
		if line != "" {
			contents = append(contents, line)
		}
	}

	result := make(map[string][]string)

	for _, content := range contents {
		contentSplit := strings.Split(content, ",")

		if len(contentSplit) != len(headerSplit) {
			return nil, errors.New("content length is not the same as header length")
		}

		for i, header := range headerSplit {
			result[header] = append(result[header], contentSplit[i])
		}
	}

	return result, nil
}
