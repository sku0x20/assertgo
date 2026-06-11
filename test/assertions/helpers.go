package assertions

import (
	"github.com/sku0x20/assertgo/pkg"
	agtest "github.com/sku0x20/assertgo/test"
)

func newSink() (*agtest.MockT, *pkg.TSink) {
	mock := &agtest.MockT{}
	return mock, pkg.NewTSink(mock)
}
