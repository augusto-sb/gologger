package gologger

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestDebug(t *testing.T) {
	Debug("hola")
	Info("hola")
	Warn("hola")
	Error("hola")
	log.Println("utc")
	fmt.Println(time.Now().Format(time.RFC3339))
	fmt.Println(time.Now().UTC().Format(time.RFC3339))
}
