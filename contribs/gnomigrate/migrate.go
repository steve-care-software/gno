package main

import (
	"github.com/steve-care-software/gno/tm2/pkg/commands"
	"github.com/steve-care-software/gnomigrate/internal/txs"
)

func newMigrateCmd(io commands.IO) *commands.Command {
	cmd := commands.NewCommand(
		commands.Metadata{
			ShortUsage: "<subcommand> [flags] [<arg>...]",
			ShortHelp:  "gno migration suite",
			LongHelp:   "Gno state migration suite, for managing legacy headaches",
		},
		commands.NewEmptyConfig(),
		commands.HelpExec,
	)

	cmd.AddSubCommands(
		txs.NewTxsCmd(io),
	)

	return cmd
}
