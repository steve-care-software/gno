package hd

import (
	"github.com/steve-care-software/gno/tm2/pkg/amino"
)

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/steve-care-software/gno/tm2/pkg/crypto/hd",
	"tm",
	amino.GetCallersDirname(),
).WithDependencies().WithTypes(
	BIP44Params{}, "Bip44Params",
))
