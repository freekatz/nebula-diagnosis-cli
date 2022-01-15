package info

import (
	"context"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/config"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/logger"
	"time"
)

func Run(conf *config.InfoConfig) {
	// TODO fix the cancel bugs
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	commonInfo := conf.Common
	for name, nodeInfo := range conf.Node {
		// FileLogger
		go func() {
			var _logger logger.Logger
			_logger = logger.GetFileLogger(name, nodeInfo.OutputDirPath)
			// the conf has been verified, so don't need to handle error
			d, _ := time.ParseDuration(nodeInfo.Duration)
			if commonInfo.Duration == "-1" {
				runWithInfinity(commonInfo, nodeInfo, _logger)
			} else if d == 0 {
				run(commonInfo, nodeInfo, _logger)
				cancel() // TODO temp code
			} else {
				runWithDuration(commonInfo, nodeInfo, _logger)
			}
		}()
		// CmdLogger
		//go func() {
		//	var _logger logger.Logger
		//	_logger = logger.GetCmdLogger(name)
		//	// the conf has been verified, so don't need to handle error
		//	d, _ := time.ParseDuration(commonInfo.Duration)
		//	if commonInfo.Duration == "-1" {
		//		runWithInfinity(commonInfo, nodeInfo, _logger)
		//	} else if d == 0 {
		//		run(commonInfo, nodeInfo, _logger)
		//		cancel() // TODO temp code
		//	} else {
		//		runWithDuration(commonInfo, nodeInfo, _logger)
		//	}
		//}()
	}

	for {
		select {
		case <-ctx.Done():
			return
		}
	}

}

func run(commonConf *config.CommonConfig, nodeInfo *config.NodeConfig, defaultLogger logger.Logger) {
	for _, option := range commonConf.Options {
		//fetchInfo(conf, option, defaultLogger)
		fetchAndSaveInfo(nodeInfo, option, defaultLogger)
	}
}

func runWithInfinity(commonConf *config.CommonConfig, nodeInfo *config.NodeConfig, defaultLogger logger.Logger) {
	p, _ := time.ParseDuration(commonConf.Period)
	ticker := time.NewTicker(p)
	for {
		select {
		case <-ticker.C:
			run(commonConf, nodeInfo, defaultLogger)
		default:

		}
	}
}

func runWithDuration(commonConf *config.CommonConfig, nodeInfo *config.NodeConfig, defaultLogger logger.Logger) {
	p, _ := time.ParseDuration(commonConf.Period)
	ticker := time.NewTicker(p)
	ch := make(chan bool)
	go func(ticker *time.Ticker) {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				run(commonConf, nodeInfo, defaultLogger)
			case stop := <-ch:
				if stop {
					return
				}
			default:

			}
		}
	}(ticker)

	d, _ := time.ParseDuration(commonConf.Duration)
	time.Sleep(d)
	ch <- true
	close(ch)
}