package main

import (
	_ "calendar-api/config"
	"calendar-api/jobs"
	"calendar-api/router"
)

func main() {
	jobs.InitSidekiq()
	router.Run()
}
