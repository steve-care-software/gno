package emitter

import "github.com/steve-care-software/gno/contribs/gnodev/pkg/events"

type NoopServer struct{}

func (*NoopServer) Emit(evt events.Event) {}
