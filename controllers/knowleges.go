package controllers

import "github.com/labstack/echo/v4"

type KnowledgeCtrl struct{}

func (ctrl KnowledgeCtrl) InitRouter(g *echo.Group) {
	g.GET("/knowledges", ctrl.Index)
}
