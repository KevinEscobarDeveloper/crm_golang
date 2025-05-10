package router

import (
	"github.com/labstack/echo/v4"
	"mango_crm/api/v1/handler"
	"mango_crm/pkg/route"
)

func Register(e *echo.Echo, h *handler.OrchardHandler) {
	e.POST(route.APIVersion, h.Create)
}
