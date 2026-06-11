package assertgo

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg"
)

func T(t testing.TB) {
	sink := pkg.NewFailSink(t)
	sink.Fail("failed")
}

// func Ts(t testing.TB) {}
