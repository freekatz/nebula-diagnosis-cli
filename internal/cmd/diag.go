package cmd

import (
	"strings"

	"github.com/1uvu/nebula-diagnosis-cli/internal/diag"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/config"
	"github.com/1uvu/nebula-diagnosis-cli/pkg/errorx"

	"github.com/urfave/cli/v2"
)

var diagCMD = &cli.Command{
	Name:  "diag",
	Usage: "diag the collected infos",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "output_dir_path",
			Aliases: []string{"O"},
			Usage:   "--output_dir_path or -O, the output dir of diag results, logs, and others output",
			Value:   "./output",
		},
		&cli.BoolFlag{
			Name:    "log_to_file",
			Aliases: []string{"L"},
			Usage:   "--log_to_file or -L, logging to file",
			Value:   false,
		},
		&cli.StringFlag{
			Name:    "input_dir_path",
			Aliases: []string{"I"},
			Usage:   "--input_dir_path or -I, the input dir of infos data",
			Value:   "",
		},
		&cli.StringFlag{
			Name:  "option",
			Usage: "the diags to analyze, included: partition, all, no, etc.",
			Value: string(config.AllDiag),
		},
	},
	Action: func(ctx *cli.Context) error {
		if GlobalDiagConfig == nil {
			GlobalDiagConfig = new(config.DiagConfig)
		}
		if ctx.IsSet("output_dir_path") {
			outputDirPath := ctx.String("output_dir_path")
			GlobalDiagConfig.OutputDirPath = outputDirPath
		}
		if ctx.IsSet("log_to_file") {
			GlobalDiagConfig.LogToFile = ctx.Bool("log_to_file")
		}
		if ctx.IsSet("input_dir_path") {
			inputDirpath := ctx.String("input_dir_path")
			GlobalDiagConfig.InputDirPath = inputDirpath
		} else {
			return errorx.ErrNoInputFilepath
		}
		if ctx.IsSet("option") {
			var options []string
			optionsStr := ctx.String("option")
			if strings.Contains(optionsStr, "all") {
				options = []string{"all"}
			} else {
				for _, optionStr := range strings.Split(optionsStr, ",") {
					options = append(options, optionStr)
				}
			}
			diagOptions := make([]config.DiagOption, len(options))
			for i := range options {
				diagOptions[i] = config.DiagOption(options[i])
			}
			GlobalDiagConfig.Options = diagOptions
		}
		GlobalDiagConfig.Complete()
		if !GlobalDiagConfig.Validate() {
			GlobalCMDLogger.Errorf("validate diag config failed.\n")
			return errorx.ErrConfigInvalid
		}
		diag.Run(GlobalDiagConfig)
		return nil
	},
}
