package controllers

import (
	"calendar-api/heartbeat"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (HeartbeatCtrl) Show(c echo.Context) error {
	hash := heartbeat.CommitHash
	if hash == "" {
		hash = heartbeat.NotAvailableMessage
	}

	uptime := time.Since(heartbeat.StartTime).String()

	return c.JSON(http.StatusOK, heartbeat.HeartbeatMessage{
		Status: "running",
		Build:  hash,
		Branch: heartbeat.GitBranch,
		Uptime: uptime})
}
