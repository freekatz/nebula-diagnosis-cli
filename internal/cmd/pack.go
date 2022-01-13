package cmd

import (
	"log"

	"github.com/1uvu/nebula-diagnosis-cli/pkg/config"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/errorx"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/utils"
	"github.com/urfave/cli/v2"
)

var packCmd = &cli.Command{
	Name:  "pack",
	Usage: "pack the collected data",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"C"},
			Usage:   "--config or -C, the config file for pack data",
			Value:   "",
		},
		&cli.StringFlag{
			Name:    "output_dir_path",
			Aliases: []string{"O"},
			Usage:   "--output_dir_path or -O, the output dir of pack results, logs, and others output",
			Value:   "./output",
		},
		&cli.StringFlag{
			Name:    "input_dir_path",
			Aliases: []string{"I"},
			Usage:   "--input_dir_path or -I, the input dir of infos data",
			Value:   "",
		},
	},
	Action: func(ctx *cli.Context) error {
		configPath := ctx.String("config")
		var err error
		if configPath != "" {
			GlobalPackConfig, err = config.NewPackConfig(configPath, utils.GetConfigType(configPath))
			if err != nil {
				return err
			}
		}
		if GlobalPackConfig == nil {
			GlobalPackConfig = new(config.PackConfig)
		}

		if ctx.IsSet("output_dir_path") {
			outputDirPath := ctx.String("output_dir_path")
			GlobalPackConfig.OutputDirPath = outputDirPath
		}
		if ctx.IsSet("input_dir_path") {
			inputDirPath := ctx.String("input_dir_path")
			GlobalPackConfig.InputDirPath = inputDirPath
		} else {
			return errorx.ErrNoInputDir
		}

		GlobalPackConfig.Complete()
		//pack.Run(GlobalPackConfig)
		log.Println(GlobalPackConfig)
		return nil
	},
}
