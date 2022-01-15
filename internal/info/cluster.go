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

	for name := range conf.Node {
		// FileLogger
		go func(name string) {
			nodeConfig := conf.Node[name]
			var _logger logger.Logger
			_logger = logger.GetFileLogger(name, nodeConfig.OutputDirPath)
			// the conf has been verified, so don't need to handle error
			d, _ := time.ParseDuration(nodeConfig.Duration)
			if nodeConfig.Duration == "-1" {
				runWithInfinity(nodeConfig, _logger)
			} else if d == 0 {
				run(nodeConfig, _logger)
				cancel() // TODO temp code
			} else {
				runWithDuration(nodeConfig, _logger)
			}
		}(name)
	}

	for {
		select {
		case <-ctx.Done():
			return
		}
	}

}

func run(nodeConfig *config.NodeConfig, defaultLogger logger.Logger) {
	for _, option := range nodeConfig.Options {
		//fetchInfo(conf, option, defaultLogger)
		fetchAndSaveInfo(nodeConfig, option, defaultLogger)
	}
}

func runWithInfinity(nodeConfig *config.NodeConfig, defaultLogger logger.Logger) {
	p, _ := time.ParseDuration(nodeConfig.Period)
	ticker := time.NewTicker(p)
	for {
		select {
		case <-ticker.C:
			run(nodeConfig, defaultLogger)
		default:

		}
	}
}

func runWithDuration(nodeConfig *config.NodeConfig, defaultLogger logger.Logger) {
	p, _ := time.ParseDuration(nodeConfig.Period)
	ticker := time.NewTicker(p)
	ch := make(chan bool)
	go func(ticker *time.Ticker) {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				run(nodeConfig, defaultLogger)
			case stop := <-ch:
				if stop {
					return
				}
			default:

			}
		}
	}(ticker)

	d, _ := time.ParseDuration(nodeConfig.Duration)
	time.Sleep(d)
	ch <- true
	close(ch)
}
