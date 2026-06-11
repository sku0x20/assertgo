package assertions

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/assertions"
	agtest "github.com/sku0x20/assertgo/test"
)

func Test_AnyAssertions_IsEqualTo_pass(t *testing.T) {
	mock, sink := agtest.NewSink()
	a := assertions.NewAnyAssertions(sink, "hello")
	a.IsEqualTo("hello")
	if mock.FatalCalled {
		t.Fatal("expected no failure")
	}
}

func Test_AnyAssertions_IsEqualTo_fail(t *testing.T) {
	mock, sink := agtest.NewSink()
	a := assertions.NewAnyAssertions(sink, "hello")
	a.IsEqualTo("world")
	if !mock.FatalCalled {
		t.Fatal("expected failure")
	}
}
