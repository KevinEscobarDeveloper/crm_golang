package handler

import (
	"github.com/labstack/echo/v4"
	"mango_crm/api/v1/request"
	"mango_crm/internal/orchard/domain/usecase"
	"net/http"
)

type OrchardHandler struct {
	usecase usecase.OrchardUseCase
}

func NewOrchardHandler(u usecase.OrchardUseCase) *OrchardHandler {
	return &OrchardHandler{usecase: u}
}

func (h *OrchardHandler) Create(c echo.Context) error {
	var req request.CreateOrchardRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"id": "1"})
}
