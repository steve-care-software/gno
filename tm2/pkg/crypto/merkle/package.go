package merkle

import (
	"github.com/steve-care-software/gno/tm2/pkg/amino"
)

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/steve-care-software/gno/tm2/pkg/crypto/merkle",
	"tm",
	amino.GetCallersDirname(),
).WithDependencies().WithTypes(
	ProofOp{},
	Proof{},
	SimpleProof{},
	SimpleProofNode{},
))
