package main

import (
	"calendar-api/jobs"

	_ "calendar-api/config"

	"github.com/jrallison/go-workers"
)

func process() {
	workers.Process(jobs.CrmLeadsQueue, jobs.AddCrmLeads, 10)
}

func main() {
	jobs.InitSidekiq()

	process()

	go workers.StatsServer(7001)

	workers.Run()
}
