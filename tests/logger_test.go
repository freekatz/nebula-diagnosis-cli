package tests

import (
	"testing"

	"github.com/1uvu/nebula-diagnosis-cli/pkg/logger"
)

func TestGetCMDLogger(t *testing.T) {
	cmdLogger := logger.GetLogger("test_cmd", "", false)
	cmdLogger.Info("test cmd logger")
}

func TestGetFileLogger(t *testing.T) {
	fileLogger := logger.GetLogger("test_file", "./tmp", true)
	fileLogger.Info("test cmd logger")
}
