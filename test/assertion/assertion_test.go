package assertion

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/assertion"
	"github.com/sku0x20/assertgo/pkg/matcherasserter"
	"github.com/sku0x20/assertgo/pkg/sink"
	agtest "github.com/sku0x20/assertgo/test"
)

func Test_AnyAssertion_EqualTo(t *testing.T) {
	assertPasses(t, "hello", func(a *assertion.Assertion[any]) { a.EqualTo("hello") })
}

func Test_AnyAssertion_SameAs(t *testing.T) {
	obj := new(int)
	assertPasses(t, obj, func(a *assertion.Assertion[any]) { a.SameAs(obj) })
}

func Test_AnyAssertion_IsNil(t *testing.T) {
	assertPasses(t, nil, func(a *assertion.Assertion[any]) { a.IsNil() })
}

func Test_AnyAssertion_Not(t *testing.T) {
	assert(t, "hello", true, func(a *assertion.Assertion[any]) { a.Not().EqualTo("hello") })
}

func newAssertion(s *sink.TSink, value any) *assertion.Assertion[any] {
	return assertion.New(matcherasserter.New(s, value))
}

func assertPasses(t *testing.T, value any, fn func(*assertion.Assertion[any])) {
	assert(t, value, false, fn)
}

func assert(t *testing.T, value any, fail bool, fn func(*assertion.Assertion[any])) {
	t.Helper()
	mock, s := agtest.NewSink()
	fn(newAssertion(s, value))
	if mock.FatalCalled != fail {
		t.Fatal("expected failure")
	}
}
