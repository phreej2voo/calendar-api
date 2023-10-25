package heartbeat

import (
	"time"
)

const NotAvailableMessage = "Not available"

var CommitHash string
var StartTime time.Time
var GitBranch string

type HeartbeatMessage struct {
	Status string `json:"status"`
	Build  string `json:"build"`
	Branch string `json:"branch"`
	Uptime string `json:"uptime"`
}

func init() {
	StartTime = time.Now()
}
