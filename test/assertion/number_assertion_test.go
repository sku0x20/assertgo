package assertion

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/assertion"
	"github.com/sku0x20/assertgo/pkg/matcherasserter"
	agtest "github.com/sku0x20/assertgo/test"
)

func Test_NumberAssertion_GreaterThan(t *testing.T) {
	assertNumberPasses(t, 10, func(a *assertion.NumberAssertion[int]) { a.GreaterThan(5) })
}

func Test_NumberAssertion_GreaterThanOrEqual(t *testing.T) {
	assertNumberPasses(t, 10, func(a *assertion.NumberAssertion[int]) { a.GreaterThanOrEqual(10) })
}

func Test_NumberAssertion_LessThan(t *testing.T) {
	assertNumberPasses(t, 5, func(a *assertion.NumberAssertion[int]) { a.LessThan(10) })
}

func Test_NumberAssertion_LessThanOrEqual(t *testing.T) {
	assertNumberPasses(t, 10, func(a *assertion.NumberAssertion[int]) { a.LessThanOrEqual(10) })
}

func Test_NumberAssertion_InRange(t *testing.T) {
	assertNumberPasses(t, 5, func(a *assertion.NumberAssertion[int]) { a.InRange(1, 10) })
}

func Test_NumberAssertion_Not(t *testing.T) {
	assertNumberPasses(t, 3, func(a *assertion.NumberAssertion[int]) { a.Not().GreaterThan(5) })
}

func assertNumberPasses(t *testing.T, value int, fn func(*assertion.NumberAssertion[int])) {
	t.Helper()
	mock, s := agtest.NewSink()
	a := assertion.NewNumber[int](matcherasserter.New(s, value))
	fn(a)
	if mock.FatalCalled {
		t.Fatal("expected no failure")
	}
}
