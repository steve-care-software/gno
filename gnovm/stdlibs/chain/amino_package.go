package chain

import (
	"github.com/steve-care-software/gno/tm2/pkg/amino"
	abci "github.com/steve-care-software/gno/tm2/pkg/bft/abci/types"
)

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/steve-care-software/gno/gnovm/stdlibs/chain",
	"tm",
	amino.GetCallersDirname(),
).
	WithDependencies(
		abci.Package,
	).
	WithTypes(
		EventAttribute{},
		Event{},
		StorageDepositEvent{},
		StorageUnlockEvent{},
	))
