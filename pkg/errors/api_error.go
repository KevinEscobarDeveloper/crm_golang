package errors

import (
	"mango_crm/api/v1/response"
	"net/http"
)

func NewApiError(statusCode int, message interface{}) *response.ApiError {
	return &response.ApiError{
		StatusCode: statusCode,
		Message:    message,
	}
}

func InvalidRequestFormat() *response.ApiError {
	return NewApiError(http.StatusBadRequest, "Invalid request format")
}

func ValidationError(validationErrors map[string]string) *response.ApiError {
	return NewApiError(http.StatusUnprocessableEntity, validationErrors)
}

func NotFoundError(message string) *response.ApiError {
	return NewApiError(http.StatusNotFound, message)
}
