package logger

import (
	"testing"
)

func TestInfo(t *testing.T) {
	Info("normal info")
	Fatal("file fatal")
	Error("server error")
	Debug("dev debug")
}
