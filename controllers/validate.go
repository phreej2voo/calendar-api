package controllers

import (
	"github.com/labstack/echo/v4"
)

func BindValidate(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		return NewCustomError(ParamError, err.Error())
	}
	if err := c.Validate(i); err != nil {
		return NewCustomError(ParamError, err.Error())
	}
	return nil
}
