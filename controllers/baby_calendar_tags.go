package controllers

import "github.com/labstack/echo/v4"

type BabyCalendarTagCtrl struct{}

func (ctrl BabyCalendarTagCtrl) InitRouter(g *echo.Group) {
	g.POST("/baby_calendar_tags", ctrl.Create)
	g.GET("/baby_calendar_tags/:id", ctrl.Show)
	g.POST("/baby_calendar_tags/:id/photos", ctrl.AddPhoto)
	g.GET("/baby_calendar_tags/:id/share", ctrl.Share)
}
