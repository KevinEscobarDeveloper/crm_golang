package router

import (
	"github.com/labstack/echo/v4"
	"mango_crm/api/v1/handler"
	"mango_crm/api/v1/request"
	"mango_crm/pkg/constants"
	"mango_crm/pkg/middleware"
)

func Register(e *echo.Echo, h *handler.OrchardHandler) {
	g := e.Group(constants.APIVersion)
	g.POST(constants.OrchardBasePath, h.Create, middleware.ValidationMiddlware(request.CreateOrchardRequest{}))
}
