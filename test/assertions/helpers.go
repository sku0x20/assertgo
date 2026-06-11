package assertions

import (
	"github.com/sku0x20/assertgo/pkg/sink"
	agtest "github.com/sku0x20/assertgo/test"
)

func newSink() (*agtest.MockT, *sink.TSink) {
	mock := &agtest.MockT{}
	return mock, sink.NewTSink(mock)
}
