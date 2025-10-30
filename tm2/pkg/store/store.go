package store

import (
	dbm "github.com/steve-care-software/gno/tm2/pkg/db"

	"github.com/steve-care-software/gno/tm2/pkg/store/rootmulti"
	"github.com/steve-care-software/gno/tm2/pkg/store/types"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	return rootmulti.NewMultiStore(db)
}
