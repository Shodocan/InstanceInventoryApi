package logger

import (
	"fmt"
	"log"

	"github.com/Shodocan/InstanceInventoryApi/configs"
)

var (
	Debug string = "Debug"
	Info  string = "Info"
)

func ErrorIf(format string, err error, args ...interface{}) {
	if err != nil {
		args = append(args, err)
		Errorf(format, args...)
	}
}

func Debugf(format string, args ...interface{}) {
	if configs.LogLevel == Debug {
		fmt.Printf(format+"\n", args...)
	}
}
func Infof(format string, args ...interface{}) {
	if configs.LogLevel == Debug || configs.LogLevel == Info {
		fmt.Printf(format+"\n", args...)
	}
}
func Errorf(format string, args ...interface{}) {
	log.Printf(format+"\n", args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Panicf(format+"\n", args...)
}
