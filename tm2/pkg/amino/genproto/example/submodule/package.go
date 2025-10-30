package submodule

import (
	"github.com/steve-care-software/gno/tm2/pkg/amino"
	"github.com/steve-care-software/gno/tm2/pkg/amino/genproto/example/submodule2"
)

var Package = amino.RegisterPackage(
	amino.NewPackage(
		"github.com/steve-care-software/gno/tm2/pkg/amino/genproto/example/submodule",
		"submodule",
		amino.GetCallersDirname(),
	).WithDependencies(
		submodule2.Package,
	).WithTypes(
		StructSM{},
	),
)
