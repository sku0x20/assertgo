package assertion

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/assertion"
	"github.com/sku0x20/assertgo/pkg/matcher"
	"github.com/sku0x20/assertgo/pkg/matcherasserter"
	agtest "github.com/sku0x20/assertgo/test"
)

func Test_Assertion_EqualTo(t *testing.T) {
	assertPasses(t, "hello", func(a *assertion.Assertion[any]) { a.EqualTo("hello") })
}

func Test_Assertion_SameAs(t *testing.T) {
	obj := new(int)
	assertPasses(t, obj, func(a *assertion.Assertion[any]) { a.SameAs(obj) })
}

func Test_Assertion_IsNil(t *testing.T) {
	assertPasses(t, nil, func(a *assertion.Assertion[any]) { a.IsNil() })
}

func Test_Assertion_Matches(t *testing.T) {
	assertPasses(t, "hello", func(a *assertion.Assertion[any]) { a.Matches(matcher.NewDeepEqualMatcher[any]("hello")) })
}

func Test_Assertion_Not(t *testing.T) {
	assert(t, "hello", true, func(a *assertion.Assertion[any]) { a.Not().EqualTo("hello") })
}

func assertPasses(t *testing.T, value any, fn func(*assertion.Assertion[any])) {
	assert(t, value, false, fn)
}

func assert(t *testing.T, value any, fail bool, fn func(*assertion.Assertion[any])) {
	t.Helper()
	mock, s := agtest.NewSink()
	a := assertion.New(matcherasserter.New(s, value))
	fn(a)
	if mock.FatalCalled != fail {
		t.Fatal("expected failure")
	}
}
