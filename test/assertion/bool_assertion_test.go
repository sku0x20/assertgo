package assertion

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/assertion"
	"github.com/sku0x20/assertgo/pkg/matcherasserter"
	agtest "github.com/sku0x20/assertgo/test"
)

func Test_BoolAssertion_IsTrue(t *testing.T) {
	assertBoolPasses(t, true, func(a *assertion.BoolAssertion) { a.IsTrue() })
}

func Test_BoolAssertion_IsFalse(t *testing.T) {
	assertBoolPasses(t, false, func(a *assertion.BoolAssertion) { a.IsFalse() })
}

func Test_BoolAssertion_Not(t *testing.T) {
	assertBoolPasses(t, false, func(a *assertion.BoolAssertion) { a.Not().IsTrue() })
}

func assertBoolPasses(t *testing.T, value bool, fn func(*assertion.BoolAssertion)) {
	t.Helper()
	mock, s := agtest.NewSink()
	a := assertion.NewBool(matcherasserter.New(s, value))
	fn(a)
	if mock.FatalCalled {
		t.Fatal("expected no failure")
	}
}
