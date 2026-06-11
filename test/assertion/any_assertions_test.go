package assertion

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/assertion"
	agtest "github.com/sku0x20/assertgo/test"
)

func Test_AnyAssertion_IsEqualTo(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		mock, sink := agtest.NewSink()
		a := assertion.NewAnyAssertion(sink, "hello")
		a.IsEqualTo("hello")
		if mock.FatalCalled {
			t.Fatal("expected no failure")
		}
	})
	t.Run("fail", func(t *testing.T) {
		mock, sink := agtest.NewSink()
		a := assertion.NewAnyAssertion(sink, "hello")
		a.IsEqualTo("world")
		if !mock.FatalCalled {
			t.Fatal("expected failure")
		}
	})
}

func Test_AnyAssertion_IsNotEqualTo(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		mock, sink := agtest.NewSink()
		a := assertion.NewAnyAssertion(sink, "hello")
		a.IsNotEqualTo("world")
		if mock.FatalCalled {
			t.Fatal("expected no failure")
		}
	})
	t.Run("fail", func(t *testing.T) {
		mock, sink := agtest.NewSink()
		a := assertion.NewAnyAssertion(sink, "hello")
		a.IsNotEqualTo("hello")
		if !mock.FatalCalled {
			t.Fatal("expected failure")
		}
	})
}

func Test_AnyAssertion_IsNotSameAs(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		mock, sink := agtest.NewSink()
		a := assertion.NewAnyAssertion(sink, new(int))
		a.IsNotSameAs(new(int))
		if mock.FatalCalled {
			t.Fatal("expected no failure")
		}
	})
	t.Run("fail", func(t *testing.T) {
		mock, sink := agtest.NewSink()
		obj := new(int)
		a := assertion.NewAnyAssertion(sink, obj)
		a.IsNotSameAs(obj)
		if !mock.FatalCalled {
			t.Fatal("expected failure")
		}
	})
}

func Test_AnyAssertion_IsSameAs(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		mock, sink := agtest.NewSink()
		obj := new(int)
		a := assertion.NewAnyAssertion(sink, obj)
		a.IsSameAs(obj)
		if mock.FatalCalled {
			t.Fatal("expected no failure")
		}
	})
	t.Run("fail", func(t *testing.T) {
		mock, sink := agtest.NewSink()
		a := assertion.NewAnyAssertion(sink, new(int))
		a.IsSameAs(new(int))
		if !mock.FatalCalled {
			t.Fatal("expected failure")
		}
	})
}
