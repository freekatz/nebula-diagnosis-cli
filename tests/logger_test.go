package tests

import (
	"testing"

	"github.com/1uvu/nebula-diagnosis-cli/pkg/logger"
)

func TestGetCMDLogger(t *testing.T) {
	cmdLogger := logger.GetLogger("test_cmd", "")
	cmdLogger.Info(false, "test cmd logger")
}

func TestGetFileLogger(t *testing.T) {
	fileLogger := logger.GetLogger("test_file", "./tmp")
	fileLogger.Info(true, "test cmd logger")
}
