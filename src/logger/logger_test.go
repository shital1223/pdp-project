package logger_test

import (
	"logger"
	"testing"
)

func TestInit(t *testing.T) {
	config := &logger.LoggerConfig{}
	if !logger.Init(config) {
		t.Fail()
	}
}
