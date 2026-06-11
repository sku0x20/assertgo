package assertion

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/assertion"
	"github.com/sku0x20/assertgo/pkg/matcherasserter"
	"github.com/sku0x20/assertgo/pkg/sink"
	agtest "github.com/sku0x20/assertgo/test"
)

func newAssertion(s *sink.TSink, value any) *assertion.AnyAssertion {
	return assertion.NewAnyAssertion(matcherasserter.New(s, value))
}

func Test_AnyAssertion_EqualTo(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		mock, sink := agtest.NewSink()
		newAssertion(sink, "hello").EqualTo("hello")
		if mock.FatalCalled {
			t.Fatal("expected no failure")
		}
	})
	t.Run("fail", func(t *testing.T) {
		mock, sink := agtest.NewSink()
		newAssertion(sink, "hello").EqualTo("world")
		if !mock.FatalCalled {
			t.Fatal("expected failure")
		}
	})
}

func Test_AnyAssertion_IsNil(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		mock, sink := agtest.NewSink()
		newAssertion(sink, nil).IsNil()
		if mock.FatalCalled {
			t.Fatal("expected no failure")
		}
	})
	t.Run("fail", func(t *testing.T) {
		mock, sink := agtest.NewSink()
		newAssertion(sink, "hello").IsNil()
		if !mock.FatalCalled {
			t.Fatal("expected failure")
		}
	})
}

func Test_AnyAssertion_IsNotNil(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		mock, sink := agtest.NewSink()
		newAssertion(sink, "hello").IsNotNil()
		if mock.FatalCalled {
			t.Fatal("expected no failure")
		}
	})
	t.Run("fail", func(t *testing.T) {
		mock, sink := agtest.NewSink()
		newAssertion(sink, nil).IsNotNil()
		if !mock.FatalCalled {
			t.Fatal("expected failure")
		}
	})
}

func Test_AnyAssertion_NotEqualTo(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		mock, sink := agtest.NewSink()
		newAssertion(sink, "hello").Not().EqualTo("world")
		if mock.FatalCalled {
			t.Fatal("expected no failure")
		}
	})
	t.Run("fail", func(t *testing.T) {
		mock, sink := agtest.NewSink()
		newAssertion(sink, "hello").Not().EqualTo("hello")
		if !mock.FatalCalled {
			t.Fatal("expected failure")
		}
	})
}

func Test_AnyAssertion_Not().SameAs(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		mock, sink := agtest.NewSink()
		newAssertion(sink, new(int)).Not().SameAs(new(int))
		if mock.FatalCalled {
			t.Fatal("expected no failure")
		}
	})
	t.Run("fail", func(t *testing.T) {
		mock, sink := agtest.NewSink()
		obj := new(int)
		newAssertion(sink, obj).Not().SameAs(obj)
		if !mock.FatalCalled {
			t.Fatal("expected failure")
		}
	})
}

func Test_AnyAssertion_SameAs(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		mock, sink := agtest.NewSink()
		obj := new(int)
		newAssertion(sink, obj).SameAs(obj)
		if mock.FatalCalled {
			t.Fatal("expected no failure")
		}
	})
	t.Run("fail", func(t *testing.T) {
		mock, sink := agtest.NewSink()
		newAssertion(sink, new(int)).SameAs(new(int))
		if !mock.FatalCalled {
			t.Fatal("expected failure")
		}
	})
}
