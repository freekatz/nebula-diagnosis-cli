package cmd

import (
	"github.com/nebula/nebula-diagnose/internal/pack"
	"github.com/nebula/nebula-diagnose/pkg/config"
	"github.com/nebula/nebula-diagnose/pkg/errorx"
	"github.com/nebula/nebula-diagnose/pkg/utils"

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
			Name:     "tar_filepath",
			Aliases:  []string{"I"},
			Usage:    "--tar_filepath or -I, the input tar filepath",
			Required: true,
		},
		&cli.StringFlag{
			Name:        "tar_filename",
			Aliases:     []string{"N"},
			Usage:       "--tar_filename or -N, the output tar filename, will auto complete the suffix and output in current root dir",
			DefaultText: "basename of input tar filepath",
		},
	},
	Action: func(ctx *cli.Context) error {
		var err error
		if ctx.IsSet("config") {
			configPath := ctx.String("config")
			GlobalPackConfig, err = config.NewPackConfig(configPath, utils.GetConfigType(configPath))
			if err != nil {
				GlobalCMDLogger.Errorf("has error to create pack config.\n")
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
		if ctx.IsSet("tar_filepath") {
			tarFilepath := ctx.String("tar_filepath")
			GlobalPackConfig.TarFilepath = tarFilepath
		} else {
			if GlobalPackConfig.TarFilepath == "" {
				return errorx.ErrNoInputFilepath
			}
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
