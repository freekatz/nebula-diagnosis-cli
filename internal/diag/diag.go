package diag

import (
	"github.com/nebula/nebula-diagnose/pkg/config"
	"github.com/nebula/nebula-diagnose/pkg/logger"
)

func Run(conf *config.DiagConfig) {
	var _logger logger.Logger
	_logger = logger.GetLogger("diag", conf.OutputDirPath, conf.LogToFile)
	_logger.Info(conf)
	_logger.Info(conf)
}
