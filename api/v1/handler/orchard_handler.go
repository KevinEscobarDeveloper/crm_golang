package handler

import (
	"github.com/labstack/echo/v4"
	"mango_crm/api/v1/request"
	"mango_crm/internal/orchard/domain/usecase"
	"mango_crm/pkg/logger"
	"net/http"
)

type OrchardHandler struct {
	usecase usecase.OrchardUseCase
	logger  *logger.Logger
}

func NewOrchardHandler(u usecase.OrchardUseCase, log *logger.Logger) *OrchardHandler {
	return &OrchardHandler{usecase: u, logger: log}
}

func (h *OrchardHandler) Create(c echo.Context) error {
	h.logger.Info("Creating new orchard")
	_ = c.Get("body").(*request.CreateOrchardRequest)

	return c.JSON(http.StatusCreated, map[string]string{"id": "1"})
}
