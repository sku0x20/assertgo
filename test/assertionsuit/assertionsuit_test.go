package assertionsuit

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/assertion"
	"github.com/sku0x20/assertgo/pkg/assertionsuit"
	agtest "github.com/sku0x20/assertgo/test"
)

func Test_AssertionSuit_Assert_returnsAssertion(t *testing.T) {
	assertType[*assertion.Assertion[any]](t, newSuit().Assert("hello"))
}

func Test_AssertionSuit_AssertBool_returnsBoolAssertion(t *testing.T) {
	assertType[*assertion.BoolAssertion](t, newSuit().AssertBool(true))
}

func Test_AssertionSuit_AssertNumber_returnsNumberAssertion(t *testing.T) {
	assertType[*assertion.NumberAssertion[int]](t, newSuit().AssertNumber(10))
}

func Test_AssertionSuit_AssertSlice_returnsSliceAssertion(t *testing.T) {
	assertType[*assertion.SliceAssertion[any]](t, newSuit().AssertSlice([]any{1, 2, 3}))
}

func newSuit() *assertionsuit.AssertionSuit {
	_, s := agtest.NewSink()
	return assertionsuit.New(s)
}

func assertType[T any](t *testing.T, value any) {
	t.Helper()
	if _, ok := value.(T); !ok {
		t.Fatalf("expected %T", *new(T))
	}
}
