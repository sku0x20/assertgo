package assertgo

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg"
	"github.com/sku0x20/assertgo/pkg/sink"
)

func T(t testing.TB) *pkg.AssertionSuite {
	s := sink.NewTSink(t)
	return pkg.NewAssertionSuite(s)
}

// func Ts(t testing.TB) {}
