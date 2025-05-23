package gologger

import (
	"fmt"
	//"log"
	"os"
	"time"
)

const (
	DebugLevel string = "debug"
	InfoLevel  string = "info"
	WarnLevel  string = "warn"
	ErrorLevel string = "error"
)

var loc *time.Location

func println(str string) {
	//log.Println(str)
	//ts := time.Now().In(loc).Format(time.RFC3339)
	ts := time.Now().In(loc).Format(time.DateTime)
	fmt.Println(ts + " " + str)
}

func init() {
	var err error
	//time.Local = time.UTC // utc

	var loggerColor string = os.Getenv("LOGGER_COLOR")
	var loggerLevel string = os.Getenv("LOGGER_LEVEL")
	var loggerTZ string = os.Getenv("LOGGER_TZ")

	loc, err = time.LoadLocation(loggerTZ)
	if err != nil {
		//panic("logger: invalid timezone")
		loc = time.UTC
	}

	if loggerLevel != DebugLevel && loggerLevel != WarnLevel && loggerLevel != ErrorLevel {
		loggerLevel = "info"
	}

	if loggerColor == "true" {
		switch loggerLevel {
		case DebugLevel:
			Debug = func(str string) { println("\033[34mDEBUG\033[0m " + str) }
			fallthrough
		case InfoLevel:
			Info = func(str string) { println("\033[32mINFO\033[0m  " + str) }
			fallthrough
		case WarnLevel:
			Warn = func(str string) { println("\033[33mWARN\033[0m  " + str) }
			fallthrough
		case ErrorLevel:
			Error = func(str string) { println("\033[31mERROR\033[0m " + str) }
		}
	} else {
		switch loggerLevel {
		case DebugLevel:
			Debug = func(str string) { println("DEBUG " + str) }
			fallthrough
		case InfoLevel:
			Info = func(str string) { println("INFO  " + str) }
			fallthrough
		case WarnLevel:
			Warn = func(str string) { println("WARN  " + str) }
			fallthrough
		case ErrorLevel:
			Error = func(str string) { println("ERROR " + str) }
		}
	}
}

var Debug = func(str string) {}
var Info = func(str string) {}
var Warn = func(str string) {}
var Error = func(str string) {}
