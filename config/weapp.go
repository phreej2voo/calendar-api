package config

import (
	"os"

	"github.com/medivhzhan/weapp/v3"
	"github.com/medivhzhan/weapp/v3/auth"
	"github.com/medivhzhan/weapp/v3/logger"
)

var Weapp = struct {
	Auth *auth.Auth
	Cli  *weapp.Client
}{}

func initWeapp() {
	Weapp.Cli = weapp.NewClient(os.Getenv("MP_APPID"), os.Getenv("MP_SECRET"))
	Weapp.Cli.SetLogLevel(logger.Silent)
	Weapp.Auth = Weapp.Cli.NewAuth()
}
