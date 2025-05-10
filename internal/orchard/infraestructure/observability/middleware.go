package observability

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func metricsMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			method := c.Request().Method
			route := c.Path()

			IncRequest(ctx, method, route)

			err := next(c)

			status := c.Response().Status
			if status >= http.StatusOK && status < http.StatusBadRequest {
				IncSuccess(ctx, method, route, status)
			} else {
				IncError(ctx, method, route, status)
			}

			return err
		}
	}
}
