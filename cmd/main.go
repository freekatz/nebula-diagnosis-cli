package main

import (
	"log"
	"os"

	"github.com/nebula/nebula-diagnose/internal/cmd"
	"github.com/nebula/nebula-diagnose/pkg/errorx"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                   cmd.Name,
		Usage:                  cmd.Desc,
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
