package controllers

import "github.com/labstack/echo/v4"

type PhotosCtrl struct{}

func (ctrl PhotosCtrl) InitRouter(g *echo.Group) {
	g.DELETE("/photos/:id", ctrl.Destroy)
}
