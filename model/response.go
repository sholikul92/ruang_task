package model

type HandlerResponseSuccess struct {
	Status  string `json:"status"`
	Message any    `json:"message"`
}

type HandlerResponseError struct {
	Error string `json:"error"`
}

func CreateHandlerResponseSuccess(message any) *HandlerResponseSuccess {
	return &HandlerResponseSuccess{
		Status:  "success",
		Message: message,
	}
}

func CreateHandlerResponseError(err string) *HandlerResponseError {
	return &HandlerResponseError{Error: err}
}
