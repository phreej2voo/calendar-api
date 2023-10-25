package controllers

import "github.com/labstack/echo/v4"

type HeartbeatCtrl struct{}

func (ctrl HeartbeatCtrl) InitRouter(g *echo.Group) {
	g.GET("/heartbeat", ctrl.Show)
}
