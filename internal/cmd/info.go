package cmd

import (
	"github.com/nebula/nebula-diagnose/internal/info"
	"github.com/nebula/nebula-diagnose/pkg/config"
	"github.com/nebula/nebula-diagnose/pkg/errorx"
	"github.com/nebula/nebula-diagnose/pkg/utils"

	"github.com/urfave/cli/v2"
)

var infoCMD = &cli.Command{
	Name:  "info",
	Usage: "fetch the nebula graph infos",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "config",
			Aliases:  []string{"C"},
			Usage:    "--config or -C, the config file for fetching infos",
			Required: true,
		},
	},
	Action: func(ctx *cli.Context) error {
		configPath := ctx.String("config")
		var err error

		if !ctx.IsSet("config") {
			GlobalCMDLogger.Errorf("no input info config.\n")
			return errorx.ErrNoInputConfig
		}
		GlobalInfoConfig, err = config.NewInfoConfig(configPath, utils.GetConfigType(configPath))
		if err != nil {
			GlobalCMDLogger.Errorf("has error to create info config.\n")
			return errorx.ErrConfigInvalid
		}
		info.Run(GlobalInfoConfig)
		return nil
	},
}
