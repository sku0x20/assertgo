package assertgo

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg"
)

func T(t testing.TB) *pkg.AssertionSuite {
	sink := pkg.NewTSink(t)
	return pkg.NewAssertionSuite(sink)
}

// func Ts(t testing.TB) {}
