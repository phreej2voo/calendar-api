package config

import (
	"time"
)

func initTimezone() {
	var cstZone = time.FixedZone("CST", 8*3600)
	time.Local = cstZone
}
