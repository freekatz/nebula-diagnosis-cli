package cmd

import (
	"github.com/1uvu/nebula-diagnosis-cli/internal/pack"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/config"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/errorx"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/utils"
	"github.com/urfave/cli/v2"
)

var packCMD = &cli.Command{
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
			Name:    "tar_filepath",
			Aliases: []string{"P"},
			Usage:   "--tar_filepath or -P, the input tar filepath",
			Value:   "",
		},
		&cli.StringFlag{
			Name:    "tar_filename",
			Aliases: []string{"N"},
			Usage:   "--tar_filename or -N, the output tar filename, will auto complete the suffix",
			Value:   "",
		},
	},
	Action: func(ctx *cli.Context) error {
		var err error
		if ctx.IsSet("config") {
			configPath := ctx.String("config")
			GlobalPackConfig, err = config.NewPackConfig(configPath, utils.GetConfigType(configPath))
			if err != nil {
				GlobalCMDLogger.Errorf("has error to create pack config.\n")
				GlobalCMDLogger.Infof("now auto complete the pack config.\n")
			}
		}
		if GlobalPackConfig == nil {
			GlobalPackConfig = new(config.PackConfig)
		}

		if ctx.IsSet("output_dir_path") {
			outputDirPath := ctx.String("output_dir_path")
			GlobalPackConfig.OutputDirPath = outputDirPath
		}
		if ctx.IsSet("tar_filepath") {
			tarFilepath := ctx.String("tar_filepath")
			GlobalPackConfig.TarFilepath = tarFilepath
		} else {
			return errorx.ErrNoInputFilepath
		}
		if ctx.IsSet("tar_filename") {
			tarFilename := ctx.String("tar_filename")
			GlobalPackConfig.TarFilename = tarFilename
		}

		GlobalPackConfig.Complete()
		if !GlobalPackConfig.Validate() {
			GlobalCMDLogger.Errorf("validate pack config failed.\n")
			return errorx.ErrConfigInvalid
		}
		pack.Run(GlobalPackConfig)
		return nil
	},
}
