package cmd

import (
	"github.com/1uvu/nebula-diagnosis-cli/pkg/config"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/logger"

	"github.com/urfave/cli/v2"
)

const (
	Name    = "nebula diagnosis cli"
	Desc    = `A free and open source distributed diagnosis cli tool for nebula graph`
	Version = "v0.0.1"
)

var (
	Commands = []*cli.Command{
		infoCMD,
		diagCMD,
		packCMD,
	}

	GlobalInfoConfig *config.InfoConfig
	GlobalDiagConfig *config.DiagConfig
	GlobalPackConfig *config.PackConfig

	GlobalCMDLogger = logger.GetLogger("global_cli", "", false)

	GlobalOptions = []cli.Flag{
		// set the global option by &cli.XXXFlag{}
	}
	LoadGlobalOptions = func(ctx *cli.Context) error {
		// load the global option by ctx.XXX()
		return nil
	}
)
