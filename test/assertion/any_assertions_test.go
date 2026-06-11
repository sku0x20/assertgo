package assertion

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/assertion"
	"github.com/sku0x20/assertgo/pkg/matcherasserter"
	"github.com/sku0x20/assertgo/pkg/sink"
	agtest "github.com/sku0x20/assertgo/test"
)

func Test_AnyAssertion_EqualTo(t *testing.T) {
	assertPasses(t, "hello", func(a *assertion.AnyAssertion) { a.EqualTo("hello") })
}

func Test_AnyAssertion_SameAs(t *testing.T) {
	obj := new(int)
	assertPasses(t, obj, func(a *assertion.AnyAssertion) { a.SameAs(obj) })
}

func Test_AnyAssertion_IsNil(t *testing.T) {
	assertPasses(t, nil, func(a *assertion.AnyAssertion) { a.IsNil() })
}

func Test_AnyAssertion_Not(t *testing.T) {
	assertFails(t, "hello", func(a *assertion.AnyAssertion) { a.Not().EqualTo("hello") })
}

func newAssertion(s *sink.TSink, value any) *assertion.AnyAssertion {
	return assertion.NewAnyAssertion(matcherasserter.New(s, value))
}

func assertPasses(t *testing.T, value any, fn func(*assertion.AnyAssertion)) {
	t.Helper()
	mock, s := agtest.NewSink()
	fn(newAssertion(s, value))
	if mock.FatalCalled {
		t.Fatal("expected no failure")
	}
}

func assertFails(t *testing.T, value any, fn func(*assertion.AnyAssertion)) {
	t.Helper()
	mock, s := agtest.NewSink()
	fn(newAssertion(s, value))
	if !mock.FatalCalled {
		t.Fatal("expected failure")
	}
}
