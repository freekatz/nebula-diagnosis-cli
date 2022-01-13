package cmd

import (
	"log"

	"github.com/1uvu/nebula-diagnosis-cli/pkg/config"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/errorx"
	"github.com/urfave/cli/v2"
)

var unpackCmd = &cli.Command{
	Name:  "unpack",
	Usage: "unpack the collected data",
	Flags: []cli.Flag{
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
		if GlobalUnPackConfig == nil {
			GlobalUnPackConfig = new(config.UnPackConfig)
		}

		if ctx.IsSet("output_dir_path") {
			outputDirPath := ctx.String("output_dir_path")
			GlobalUnPackConfig.OutputDirPath = outputDirPath
		}
		if ctx.IsSet("input_dir_path") {
			inputDirPath := ctx.String("input_dir_path")
			GlobalUnPackConfig.InputDirPath = inputDirPath
		} else {
			return errorx.ErrNoInputDir
		}

		GlobalUnPackConfig.Complete()
		//unpack.Run(GlobalUnPackConfig)
		log.Println(GlobalUnPackConfig)
		return nil
	},
}
