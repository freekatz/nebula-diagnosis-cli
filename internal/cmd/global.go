package cmd

import (
	"github.com/nebula/nebula-diagnose/pkg/config"
	"github.com/nebula/nebula-diagnose/pkg/logger"

	"github.com/urfave/cli/v2"
)

const (
	Name    = "nebula diagnosis cli"
	Desc    = "A free and open source distributed diagnosis cli tool for nebula graph"
	Version = "v0.0.1"
	Banner  = "  _   _      _           _         _____  _                                  \n | \\ | |    | |         | |       |  __ \\(_)                                 \n |  \\| | ___| |__  _   _| | __ _  | |  | |_  __ _  __ _ _ __   ___  ___  ___ \n | . ` |/ _ \\ '_ \\| | | | |/ _` | | |  | | |/ _` |/ _` | '_ \\ / _ \\/ __|/ _ \\\n | |\\  |  __/ |_) | |_| | | (_| | | |__| | | (_| | (_| | | | | (_) \\__ \\  __/\n |_| \\_|\\___|_.__/ \\__,_|_|\\__,_| |_____/|_|\\__,_|\\__, |_| |_|\\___/|___/\\___|\n                                                   __/ |                     \n                                                  |___/                      "
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
