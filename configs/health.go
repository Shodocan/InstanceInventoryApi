package configs

import (
	"os"
	"time"
)

func GetLivingTime() time.Duration {
	duration, err := time.ParseDuration(os.Getenv("LIVING_TIME"))
	if err != nil {
		duration = time.Second * 20
	}
	return duration
}
