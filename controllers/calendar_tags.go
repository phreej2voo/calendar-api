package controllers

import "github.com/labstack/echo/v4"

type CalendarTagsCtrl struct{}

func (ctrl CalendarTagsCtrl) InitRouter(g *echo.Group) {
	g.GET("/calendar_tags", ctrl.Index)
	g.GET("/calendar_tags/achieved", ctrl.Achieved)
}
