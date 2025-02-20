package main

import (
	"log"
	"os"
	"github.com/artemka-debug/git-cli/src/commands/git-add-commit-push"
	"github.com/artemka-debug/git-cli/src/commands/merge-all-commits-into-one"
	"github.com/urfave/cli/v2"
)
func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			gacp.GACPCommand,
			gmacio.GMACIOCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
