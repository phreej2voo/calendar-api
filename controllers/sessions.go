package controllers

import "github.com/labstack/echo/v4"

type SessionsCtrl struct{}

func (ctrl SessionsCtrl) InitRouter(g *echo.Group) {
	g.POST("/login", ctrl.Create)
}
