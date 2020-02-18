package configs

import "os"

var (
	LogLevel string = os.Getenv("LOG_LEVEL")
)
