package assertgo

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/assertionsuit"
	"github.com/sku0x20/assertgo/pkg/sink"
)

func T(t testing.TB) *assertionsuit.AssertionSuit {
	s := sink.NewTSink(t)
	return assertionsuit.New(s)
}

// func Ts(t testing.TB) {}
