package diag

import (
	"github.com/1uvu/nebula-diagnosis-cli/pkg/config"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/logger"
)

func Run(conf *config.DiagConfig) {
	var _logger logger.Logger
	_logger = logger.GetLogger("diag", conf.OutputDirPath, conf.LogToFile)
	_logger.Info(conf)
	_logger.Info(conf)
}
