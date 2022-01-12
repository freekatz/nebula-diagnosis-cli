package tests

import (
	"github.com/1uvu/nebula-diagnosis-cli/pkg/logger"
	"testing"
)

func TestGetCMDLogger(t *testing.T) {
	cmdLogger := logger.GetCmdLogger("test_cmd")
	cmdLogger.Info("test cmd logger")
}

func TestGetFileLogger(t *testing.T) {
	fileLogger := logger.GetFileLogger("test_file", "./tmp")
	fileLogger.Info("test cmd logger")
}
