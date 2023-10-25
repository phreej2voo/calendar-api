package jobs

import (
	"os"

	"github.com/jrallison/go-workers"
)

func InitSidekiq() {
	workers.Configure(map[string]string{
		"server":    os.Getenv("REDIS_HOST"),
		"database":  os.Getenv("REDIS_DB"),
		"pool":      os.Getenv("REDIS_POOL"),
		"namespace": os.Getenv("REDIS_NAMESPACE"),
		"password":  os.Getenv("REDIS_PASSWORD"),
		"process":   "calendergo-workers",
	})
}
