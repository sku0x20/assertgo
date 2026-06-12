package assertion

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/assertion"
	"github.com/sku0x20/assertgo/pkg/matcherasserter"
	agtest "github.com/sku0x20/assertgo/test"
)

func Test_SliceAssertion_HasLength(t *testing.T) {
	assertSlicePasses(t, []any{1, 2, 3}, func(a *assertion.SliceAssertion[any]) {
		a.HasLength(3)
	})
}

func Test_SliceAssertion_IsEmpty(t *testing.T) {
	assertSlicePasses(t, []any{}, func(a *assertion.SliceAssertion[any]) {
		a.IsEmpty()
	})
}

func Test_SliceAssertion_EqualTo(t *testing.T) {
	assertSlicePasses(t, []any{1, 2, 3}, func(a *assertion.SliceAssertion[any]) {
		a.EqualTo([]any{1, 2, 3})
	})
}

func Test_SliceAssertion_ContainsAll(t *testing.T) {
	assertSlicePasses(t, []any{1, 2, 3}, func(a *assertion.SliceAssertion[any]) {
		a.ContainsAll([]any{1, 3})
	})
}

func Test_SliceAssertion_ContainsNone(t *testing.T) {
	assertSlicePasses(t, []any{1, 2, 3}, func(a *assertion.SliceAssertion[any]) {
		a.ContainsNone([]any{4, 5})
	})
}

func assertSlicePasses(t *testing.T, value []any, fn func(*assertion.SliceAssertion[any])) {
	t.Helper()
	mock, s := agtest.NewSink()
	a := assertion.NewSlice[any](matcherasserter.New(s, value))
	fn(a)
	if mock.FatalCalled {
		t.Fatal("expected no failure")
	}
}
