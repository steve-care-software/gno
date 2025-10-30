package gnoland

import (
	"github.com/steve-care-software/gno/tm2/pkg/amino"
)

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/steve-care-software/gno/gno.land/pkg/gnoland",
	"gno",
	amino.GetCallersDirname(),
).WithDependencies().WithTypes(
	&GnoAccount{}, "Account",
	GnoGenesisState{}, "GenesisState",
	TxWithMetadata{}, "TxWithMetadata",
	GnoTxMetadata{}, "GnoTxMetadata",
))
