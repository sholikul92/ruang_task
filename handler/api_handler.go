package handler

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"

	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type apiHandler struct {
	AiService   service.AiServiceInterface
	FileService service.FileServiceInterface
}

func NewHandler(aiService service.AiServiceInterface, fileService service.FileServiceInterface) *apiHandler {
	return &apiHandler{
		AiService:   aiService,
		FileService: fileService,
	}
}

func (h *apiHandler) ChatWithAI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload struct {
			Query string `json:"query"`
		}

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, model.CreateHandlerResponseError(err.Error()))
			return
		}

		res, err := h.AiService.GeneratedText("Qwen/Qwen2.5-1.5B-Instruct", payload.Query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.CreateHandlerResponseError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, model.CreateHandlerResponseSuccess(res))
	}
}

func (h *apiHandler) AnalyzeDataTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Request.FormValue("query")
		if query == "" {
			c.JSON(http.StatusBadRequest, model.CreateHandlerResponseError(errors.New("Query cannot be empty").Error()))
			return
		}

		file, _, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, model.CreateHandlerResponseError(err.Error()))
			return
		}
		defer file.Close()

		fileContent, err := io.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.CreateHandlerResponseError(err.Error()))
			return
		}

		table, err := h.FileService.ProcessFile(string(fileContent))
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.CreateHandlerResponseError(err.Error()))
			return
		}

		res, err := h.AiService.AnalyzeData("google/tapas-base-finetuned-wtq", table, query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.CreateHandlerResponseError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, model.CreateHandlerResponseSuccess(string(res)))
	}
}
