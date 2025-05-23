package gologger

import "testing"

func TestDebug(t *testing.T){
	Debug("hola")
	Info("hola")
	Warn("hola")
	Error("hola")
}