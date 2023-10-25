package controllers

import "github.com/labstack/echo/v4"

type BabiesCtrl struct{}

func (ctrl BabiesCtrl) InitRouter(g *echo.Group) {
	g.POST("/babies", ctrl.Create)
	g.GET("/babies", ctrl.Show)
	g.PUT("/babies", ctrl.Update)
}
