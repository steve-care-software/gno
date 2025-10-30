package main

import (
	"context"
	"os"

	"github.com/steve-care-software/gno/tm2/pkg/commands"
)

func main() {
	cmd := newGenesisCmd(commands.NewDefaultIO())

	cmd.Execute(context.Background(), os.Args[1:])
}
