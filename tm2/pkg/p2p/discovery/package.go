package discovery

import (
	"github.com/steve-care-software/gno/tm2/pkg/amino"
)

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/steve-care-software/gno/tm2/pkg/p2p/discovery",
	"p2p",
	amino.GetCallersDirname(),
).
	WithTypes(
		&Request{},
		&Response{},
	),
)
