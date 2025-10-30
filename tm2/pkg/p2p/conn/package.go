package conn

import (
	"reflect"

	"github.com/steve-care-software/gno/tm2/pkg/amino"
	"github.com/steve-care-software/gno/tm2/pkg/amino/pkg"
)

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/steve-care-software/gno/tm2/pkg/p2p/conn",
	"p2p", // keep short, do not change.
	amino.GetCallersDirname(),
).
	WithDependencies(
	// NA
	).
	WithTypes(

		// NOTE: Keep the names short.
		pkg.Type{
			Type: reflect.TypeOf(PacketPing{}),
			Name: "Ping",
		},
		pkg.Type{
			Type: reflect.TypeOf(PacketPong{}),
			Name: "Pong",
		},
		pkg.Type{
			Type: reflect.TypeOf(PacketMsg{}),
			Name: "Msg",
		},
	))
