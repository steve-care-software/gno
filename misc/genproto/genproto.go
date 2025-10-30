package main

import (
	"context"
	"os"

	"github.com/steve-care-software/gno/tm2/pkg/amino"
	"github.com/steve-care-software/gno/tm2/pkg/amino/genproto"
	"github.com/steve-care-software/gno/tm2/pkg/amino/tests"
	"github.com/steve-care-software/gno/tm2/pkg/commands"

	// TODO: move these out.
	"github.com/steve-care-software/gno/gno.land/pkg/sdk/vm"
	gno "github.com/steve-care-software/gno/gnovm/pkg/gnolang"
	abci "github.com/steve-care-software/gno/tm2/pkg/bft/abci/types"
	"github.com/steve-care-software/gno/tm2/pkg/bft/blockchain"
	"github.com/steve-care-software/gno/tm2/pkg/bft/consensus"
	ctypes "github.com/steve-care-software/gno/tm2/pkg/bft/consensus/types"
	"github.com/steve-care-software/gno/tm2/pkg/bft/mempool"
	btypes "github.com/steve-care-software/gno/tm2/pkg/bft/types"
	"github.com/steve-care-software/gno/tm2/pkg/bitarray"
	"github.com/steve-care-software/gno/tm2/pkg/crypto/ed25519"
	"github.com/steve-care-software/gno/tm2/pkg/crypto/hd"
	"github.com/steve-care-software/gno/tm2/pkg/crypto/merkle"
	"github.com/steve-care-software/gno/tm2/pkg/crypto/multisig"
	"github.com/steve-care-software/gno/tm2/pkg/sdk"
	"github.com/steve-care-software/gno/tm2/pkg/sdk/bank"
	"github.com/steve-care-software/gno/tm2/pkg/std"
)

func main() {
	cmd := commands.NewCommand(
		commands.Metadata{
			LongHelp: "Generates proto bindings for Amino packages",
		},
		commands.NewEmptyConfig(),
		execGen,
	)

	cmd.Execute(context.Background(), os.Args[1:])
}

func execGen(_ context.Context, _ []string) error {
	pkgs := []*amino.Package{
		bitarray.Package,
		merkle.Package,
		abci.Package,
		btypes.Package,
		consensus.Package,
		ctypes.Package,
		mempool.Package,
		ed25519.Package,
		blockchain.Package,
		hd.Package,
		multisig.Package,
		std.Package,
		sdk.Package,
		bank.Package,
		vm.Package,
		gno.Package,
		tests.Package,
	}

	for _, pkg := range pkgs {
		genproto.WriteProto3Schema(pkg)
		genproto.WriteProtoBindings(pkg)
		genproto.MakeProtoFolder(pkg, "proto")
	}

	for _, pkg := range pkgs {
		genproto.RunProtoc(pkg, "proto")
	}

	return nil
}
