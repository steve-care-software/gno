package main

import (
	"context"
	"os"

	"github.com/steve-care-software/gno/gno.land/pkg/keyscli"
	"github.com/steve-care-software/gno/gnovm/pkg/gnoenv"
	"github.com/steve-care-software/gno/tm2/pkg/commands"
	"github.com/steve-care-software/gno/tm2/pkg/crypto/keys/client"
)

func main() {
	baseCfg := client.DefaultBaseOptions
	baseCfg.Home = gnoenv.HomeDir()

	cmd := keyscli.NewRootCmd(commands.NewDefaultIO(), baseCfg)
	cmd.Execute(context.Background(), os.Args[1:])
}
