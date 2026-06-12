package assertionsuit

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/assertion"
	"github.com/sku0x20/assertgo/pkg/assertionsuit"
	agtest "github.com/sku0x20/assertgo/test"
)

func Test_AssertionSuit_Assert_returnsAssertion(t *testing.T) {
	var result any = newSuit(t).Assert("hello")
	if _, ok := result.(*assertion.Assertion[any]); !ok {
		t.Fatal("expected *assertion.Assertion[any]")
	}
}

func Test_AssertionSuit_AssertBool_returnsBoolAssertion(t *testing.T) {
	var result any = newSuit(t).AssertBool(true)
	if _, ok := result.(*assertion.BoolAssertion); !ok {
		t.Fatal("expected *assertion.BoolAssertion")
	}
}

func Test_AssertionSuit_AssertSlice_returnsSliceAssertion(t *testing.T) {
	var result any = newSuit(t).AssertSlice([]any{1, 2, 3})
	if _, ok := result.(*assertion.SliceAssertion[any]); !ok {
		t.Fatal("expected *assertion.SliceAssertion[any]")
	}
}

func newSuit(t *testing.T) *assertionsuit.AssertionSuit {
	_, s := agtest.NewSink()
	return assertionsuit.New(s)
}
