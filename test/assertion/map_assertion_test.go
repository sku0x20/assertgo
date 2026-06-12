package assertion

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/assertion"
	"github.com/sku0x20/assertgo/pkg/matcherasserter"
	agtest "github.com/sku0x20/assertgo/test"
)

func Test_MapAssertion_EqualTo(t *testing.T) {
	assertMapPasses(t, map[string]int{"a": 1}, func(a *assertion.MapAssertion[string, int]) {
		a.EqualTo(map[string]int{"a": 1})
	})
}

func Test_MapAssertion_ContainsAllEntries(t *testing.T) {
	assertMapPasses(t, map[string]int{"a": 1, "b": 2}, func(a *assertion.MapAssertion[string, int]) {
		a.ContainsAllEntries(map[string]int{"a": 1})
	})
}

func Test_MapAssertion_HasLength(t *testing.T) {
	assertMapPasses(t, map[string]int{"a": 1, "b": 2}, func(a *assertion.MapAssertion[string, int]) {
		a.HasLength(2)
	})
}

func Test_MapAssertion_IsEmpty(t *testing.T) {
	assertMapPasses(t, map[string]int{}, func(a *assertion.MapAssertion[string, int]) {
		a.IsEmpty()
	})
}

func Test_MapAssertion_Not(t *testing.T) {
	assertMapPasses(t, map[string]int{"a": 1}, func(a *assertion.MapAssertion[string, int]) {
		a.Not().EqualTo(map[string]int{"a": 2})
	})
}

func assertMapPasses(t *testing.T, value map[string]int, fn func(*assertion.MapAssertion[string, int])) {
	t.Helper()
	mock, s := agtest.NewSink()
	a := assertion.NewMap[string, int](matcherasserter.New(s, value))
	fn(a)
	if mock.FatalCalled {
		t.Fatal("expected no failure")
	}
}
