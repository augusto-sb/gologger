package gologger

import (
	"log"
	"os"
	"time"
)

const (
	DebugLevel string = "debug"
	InfoLevel  string = "info"
	WarnLevel  string = "warn"
	ErrorLevel string = "error"
)

func init() {

	time.Local = time.UTC // utc

	var loggerColor string = os.Getenv("LOGGER_COLOR")
	var loggerLevel string = os.Getenv("LOGGER_LEVEL")

	if loggerLevel != DebugLevel && loggerLevel != WarnLevel && loggerLevel != ErrorLevel {
		loggerLevel = "info"
	}

	if loggerColor == "true" {
		switch loggerLevel {
		case DebugLevel:
			Debug = func(str string) { log.Println("\033[34mDEBUG\033[0m " + str) }
			fallthrough
		case InfoLevel:
			Info = func(str string) { log.Println("\033[32mINFO\033[0m  " + str) }
			fallthrough
		case WarnLevel:
			Warn = func(str string) { log.Println("\033[33mWARN\033[0m  " + str) }
			fallthrough
		case ErrorLevel:
			Error = func(str string) { log.Println("\033[31mERROR\033[0m " + str) }
		}
	} else {
		switch loggerLevel {
		case DebugLevel:
			Debug = func(str string) { log.Println("DEBUG " + str) }
			fallthrough
		case InfoLevel:
			Info = func(str string) { log.Println("INFO  " + str) }
			fallthrough
		case WarnLevel:
			Warn = func(str string) { log.Println("WARN  " + str) }
			fallthrough
		case ErrorLevel:
			Error = func(str string) { log.Println("ERROR " + str) }
		}
	}
}

var Debug = func(str string) {}
var Info = func(str string) {}
var Warn = func(str string) {}
var Error = func(str string) {}
