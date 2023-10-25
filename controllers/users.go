package controllers

import "github.com/labstack/echo/v4"

type UsersCtrl struct{}

func (ctrl UsersCtrl) InitRouter(g *echo.Group) {
	g.PUT("/users/phone", ctrl.SetPhoneNumber)
	g.GET("/user", ctrl.Show)
}
