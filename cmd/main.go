package main

import (
	"github.com/nebula/nebula-diagnose/internal/cmd"
	"github.com/nebula/nebula-diagnose/pkg/errorx"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                   cmd.Name,
		Usage:                  strings.Join([]string{cmd.Desc, cmd.Banner}, "\n"),
		Version:                cmd.Version,
		UseShortOptionHandling: true,
		Flags:                  cmd.GlobalOptions,
		Before:                 cmd.LoadGlobalOptions,
		Commands:               cmd.Commands,
	}
	err := app.Run(os.Args)
	if err != nil && err != errorx.ErrPrintAndExit {
		log.Fatal(err)
	}
}
