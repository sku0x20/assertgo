package assertion

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/assertion"
	agtest "github.com/sku0x20/assertgo/test"
)

func Test_AnyAssertion_IsEqualTo_pass(t *testing.T) {
	mock, sink := agtest.NewSink()
	a := assertion.NewAnyAssertion(sink, "hello")
	a.IsEqualTo("hello")
	if mock.FatalCalled {
		t.Fatal("expected no failure")
	}
}

func Test_AnyAssertion_IsEqualTo_fail(t *testing.T) {
	mock, sink := agtest.NewSink()
	a := assertion.NewAnyAssertion(sink, "hello")
	a.IsEqualTo("world")
	if !mock.FatalCalled {
		t.Fatal("expected failure")
	}
}

func Test_AnyAssertion_IsSameAs_pass(t *testing.T) {
	mock, sink := agtest.NewSink()
	obj := new(int)
	a := assertion.NewAnyAssertion(sink, obj)
	a.IsSameAs(obj)
	if mock.FatalCalled {
		t.Fatal("expected no failure")
	}
}

func Test_AnyAssertion_IsSameAs_fail(t *testing.T) {
	mock, sink := agtest.NewSink()
	a := assertion.NewAnyAssertion(sink, new(int))
	a.IsSameAs(new(int))
	if !mock.FatalCalled {
		t.Fatal("expected failure")
	}
}
