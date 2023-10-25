package router

import (
	"calendar-api/controllers"
	"calendar-api/middlewares"
	"calendar-api/tool"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run() {
	e := echo.New()
	e.HTTPErrorHandler = controllers.CustomHTTPErrorHandler
	e.Validator = tool.NewValidation()

	middlewares.Use(e)

	app := e.Group("/app")
	controllers.HeartbeatCtrl{}.InitRouter(app)

	api := app.Group("/api")
	api.Use(middleware.KeyAuthWithConfig(middlewares.KeyAuthConfig))

	controllers.SessionsCtrl{}.InitRouter(api)
	controllers.UsersCtrl{}.InitRouter(api)
	controllers.BabiesCtrl{}.InitRouter(api)
	controllers.UsersCtrl{}.InitRouter(api)
	controllers.CalendarDaysCtrl{}.InitRouter(api)
	controllers.CalendarTagsCtrl{}.InitRouter(api)
	controllers.BabyCalendarTagCtrl{}.InitRouter(api)
	controllers.OssCtrl{}.InitRouter(api)
	controllers.KnowledgeCtrl{}.InitRouter(api)
	controllers.PhotosCtrl{}.InitRouter(api)

	e.Logger.Fatal(e.Start(":7000"))
}
