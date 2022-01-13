package cmd

import (
	"github.com/1uvu/nebula-diagnosis-cli/pkg/config"
	"github.com/urfave/cli/v2"
)

const (
	Name    = "nebula diagnosis cli"
	Desc    = `A free and open source distributed diagnosis cli tool for nebula graph`
	Version = "v0.0.1"
)

var (
	Commands = []*cli.Command{
		infoCmd,
		packCmd,
		unpackCmd,
	}
	GlobalInfoConfig   *config.InfoConfig
	GlobalPackConfig   *config.PackConfig
	GlobalUnPackConfig *config.UnPackConfig
	GlobalOptions      = []cli.Flag{
		// set the global option by &cli.XXXFlag{}
	}
	LoadGlobalOptions = func(ctx *cli.Context) error {
		// load the global option by ctx.XXX()
		return nil
	}
)
