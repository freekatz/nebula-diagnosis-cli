package cmd

import (
	"github.com/1uvu/nebula-diagnosis-cli/pkg/config"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/errorx"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/utils"
	"github.com/urfave/cli/v2"
	"log"
)

var infoCmd = &cli.Command{
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
		if configPath != "" {
			GlobalInfoConfig, err = config.NewInfoConfig(configPath, utils.GetConfigType(configPath))
			if err != nil {
				return err
			}
		}
		if GlobalInfoConfig == nil {
			return errorx.ErrNoConfig
		}
		//info.Run(GlobalInfoConfig)
		log.Println(GlobalInfoConfig)
		return nil
	},
}
