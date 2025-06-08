package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"mango_crm/pkg/errors"
	"reflect"
)

var validate = validator.New()

func ValidationMiddlware(targetStruct any) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			body := reflect.New(reflect.TypeOf(targetStruct)).Interface()

			if err := c.Bind(body); err != nil {
				return errors.InvalidRequestFormat()
			}

			if err := validate.Struct(body); err != nil {
				validationErrors := make(map[string]string)
				for _, fieldErr := range err.(validator.ValidationErrors) {
					validationErrors[fieldErr.Field()] = fieldErr.Tag()
				}
				return errors.ValidationError(validationErrors)
			}
			c.Set("body", body)
			return next(c)
		}
	}
}
