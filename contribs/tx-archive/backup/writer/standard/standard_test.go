package standard

import (
	"bytes"
	"testing"
	"time"

	"github.com/steve-care-software/gno/gno.land/pkg/gnoland"
	"github.com/steve-care-software/gno/tm2/pkg/amino"
	"github.com/steve-care-software/gno/tm2/pkg/std"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWriter_Standard(t *testing.T) {
	t.Parallel()

	var (
		b bytes.Buffer

		txData = &gnoland.TxWithMetadata{
			Tx: std.Tx{
				Memo: "example tx",
			},
			Metadata: &gnoland.GnoTxMetadata{
				Timestamp: time.Now().Unix(),
			},
		}
	)

	// Create a new standard writer
	w := NewWriter(&b)

	// Write example tx data
	require.NoError(t, w.WriteTxData(txData))

	var readTx gnoland.TxWithMetadata

	readErr := amino.UnmarshalJSON(b.Bytes(), &readTx)
	require.NoError(t, readErr)

	assert.Equal(t, *txData, readTx)
}
