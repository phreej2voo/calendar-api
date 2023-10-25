package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Use(e *echo.Echo) {
	e.Use(middleware.RequestID())
	e.Use(Logger)
}
