package assertgo

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg"
)

func T(t testing.TB) *pkg.Assertions {
	sink := pkg.NewTSink(t)
	assertions := pkg.NewAssertions(sink)
	return assertions
}

// func Ts(t testing.TB) {}
