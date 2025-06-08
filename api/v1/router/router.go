package router

import (
	"github.com/labstack/echo/v4"
	"mango_crm/api/v1/handler"
	"mango_crm/internal/orchard/domain/usecase"
	"mango_crm/pkg/logger"
)

type RouterConfig struct {
	Echo           *echo.Echo
	OrchardUseCase usecase.OrchardUseCase
	Log            *logger.Logger
}

func RegisterRoutes(config RouterConfig) {
	orchardHandler := handler.NewOrchardHandler(config.OrchardUseCase, config.Log)
	Register(config.Echo, orchardHandler)

}
