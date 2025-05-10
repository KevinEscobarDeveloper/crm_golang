package router

import (
	"github.com/labstack/echo/v4"
	"mango_crm/api/v1/handler"
	"mango_crm/internal/orchard/domain/usecase"
)

type RouterConfig struct {
	Echo           *echo.Echo
	OrchardUseCase usecase.OrchardUseCase
}

func RegisterRoutes(config RouterConfig) {
	orchardHandler := handler.NewOrchardHandler(config.OrchardUseCase)
	Register(config.Echo, orchardHandler)

}
