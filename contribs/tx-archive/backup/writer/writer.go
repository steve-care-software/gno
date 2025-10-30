package writer

import (
	"github.com/steve-care-software/gno/gno.land/pkg/gnoland"
)

// Writer defines the backup writer interface
type Writer interface {
	// WriteTxData outputs the given TX data
	// to some kind of storage
	WriteTxData(*gnoland.TxWithMetadata) error
}
