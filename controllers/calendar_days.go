package controllers

import "github.com/labstack/echo/v4"

type CalendarDaysCtrl struct{}

func (ctrl CalendarDaysCtrl) InitRouter(g *echo.Group) {
	g.GET("/calendar_days/landing", ctrl.Landing)
	g.GET("/calendar_day", ctrl.Show)
}
