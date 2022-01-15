package cmd

import (
	"github.com/1uvu/nebula-diagnosis-cli/internal/info"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/config"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/errorx"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/utils"
	"github.com/urfave/cli/v2"
)

var infoCMD = &cli.Command{
	Name:  "info",
	Usage: "fetch the nebula graph infos",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"C"},
			Usage:   "--config or -C, the config file for fetching infos",
			Value:   "",
		},
	},
	Action: func(ctx *cli.Context) error {
		configPath := ctx.String("config")
		var err error

		if !ctx.IsSet("config") {
			GlobalCMDLogger.Errorf(false, "no input info config.\n")
			return errorx.ErrNoInputConfig
		}
		GlobalInfoConfig, err = config.NewInfoConfig(configPath, utils.GetConfigType(configPath))
		if err != nil {
			GlobalCMDLogger.Errorf(false, "has error to create info config.\n")
			return errorx.ErrConfigInvalid
		}
		info.Run(GlobalInfoConfig)
		return nil
	},
}
