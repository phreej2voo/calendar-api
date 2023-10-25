package controllers

import "github.com/labstack/echo/v4"

type OssCtrl struct{}

func (ctrl OssCtrl) InitRouter(g *echo.Group) {
	g.GET("/oss/aliyun_policy", ctrl.getAliyunPolicy)
}
