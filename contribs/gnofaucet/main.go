package main

import (
	"context"
	"os"

	"github.com/steve-care-software/gno/tm2/pkg/commands"
)

func main() {
	cmd := commands.NewCommand(
		commands.Metadata{
			ShortUsage: "<subcommand> [flags] [<arg>...]",
			LongHelp:   "Starts the fund faucet that can be used by users",
		},
		commands.NewEmptyConfig(),
		commands.HelpExec,
	)

	cmd.AddSubCommands(
		newServeCmd(),
	)

	cmd.Execute(context.Background(), os.Args[1:])
}
