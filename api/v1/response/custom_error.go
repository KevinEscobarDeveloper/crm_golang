package response

import "fmt"

type ApiError struct {
	StatusCode int         `json:"status_code"`
	Message    interface{} `json:"message"`
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("%v", e.Message)
}
