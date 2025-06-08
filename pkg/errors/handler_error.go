package errors

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandleError(err error, c echo.Context) {
	var (
		code                = http.StatusInternalServerError
		message interface{} = "Internal Server Error"
	)

	var apiErr *ApiError
	if errors.As(err, &apiErr) {
		code = apiErr.StatusCode
		message = apiErr.Message
	} else {
		var httpErr *echo.HTTPError
		if errors.As(err, &httpErr) {
			code = httpErr.Code
			message = httpErr.Message
		}
	}

	if !c.Response().Committed {
		_ = c.JSON(code, map[string]interface{}{
			"status_code": code,
			"message":     message,
		})
	}
}
