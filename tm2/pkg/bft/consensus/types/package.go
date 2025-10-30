package cstypes

import (
	"github.com/steve-care-software/gno/tm2/pkg/amino"
	abci "github.com/steve-care-software/gno/tm2/pkg/bft/abci/types"
	btypes "github.com/steve-care-software/gno/tm2/pkg/bft/types"
)

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/steve-care-software/gno/tm2/pkg/bft/consensus/types",
	"tm",
	amino.GetCallersDirname(),
).
	WithGoPkgName("cstypes").
	WithDependencies(
		abci.Package,
		btypes.Package,
	).
	WithTypes(

		// Round state types
		&RoundState{},
		HRS{},
		RoundStateSimple{},
		PeerRoundState{},

		// Misc
		HeightVoteSet{},

		// Event types
		EventNewRoundStep{},
		EventNewValidBlock{},
		EventNewRound{},
		EventCompleteProposal{},
		EventTimeoutPropose{},
		EventTimeoutWait{},
	))
