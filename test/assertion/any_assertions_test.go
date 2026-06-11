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

func Test_AnyAssertion_EqualTo(t *testing.T) {
	assertPasses(t, "hello", func(a *assertion.AnyAssertion) { a.EqualTo("hello") })
	assertFails(t, "hello", func(a *assertion.AnyAssertion) { a.EqualTo("world") })
}

func Test_AnyAssertion_IsNil(t *testing.T) {
	assertPasses(t, nil, func(a *assertion.AnyAssertion) { a.IsNil() })
	assertFails(t, "hello", func(a *assertion.AnyAssertion) { a.IsNil() })
}

func Test_AnyAssertion_IsNotNil(t *testing.T) {
	assertPasses(t, "hello", func(a *assertion.AnyAssertion) { a.Not().IsNil() })
	assertFails(t, nil, func(a *assertion.AnyAssertion) { a.Not().IsNil() })
}

func Test_AnyAssertion_NotEqualTo(t *testing.T) {
	assertPasses(t, "hello", func(a *assertion.AnyAssertion) { a.Not().EqualTo("world") })
	assertFails(t, "hello", func(a *assertion.AnyAssertion) { a.Not().EqualTo("hello") })
}

func Test_AnyAssertion_SameAs(t *testing.T) {
	obj := new(int)
	assertPasses(t, obj, func(a *assertion.AnyAssertion) { a.SameAs(obj) })
	assertFails(t, new(int), func(a *assertion.AnyAssertion) { a.SameAs(new(int)) })
}

func Test_AnyAssertion_NotSameAs(t *testing.T) {
	obj := new(int)
	assertPasses(t, new(int), func(a *assertion.AnyAssertion) { a.Not().SameAs(new(int)) })
	assertFails(t, obj, func(a *assertion.AnyAssertion) { a.Not().SameAs(obj) })
}
