package pkgdownload

import (
	"errors"

	"github.com/steve-care-software/gno/tm2/pkg/std"
)

type PackageFetcher interface {
	FetchPackage(pkgPath string) ([]*std.MemFile, error)
}

func NewNoopFetcher() PackageFetcher {
	return &noopFetcher{}
}

type noopFetcher struct{}

func (nf *noopFetcher) FetchPackage(pkgPath string) ([]*std.MemFile, error) {
	return nil, errors.New("noop")
}
