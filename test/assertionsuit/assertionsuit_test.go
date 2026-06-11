package assertionsuit

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/assertion"
	"github.com/sku0x20/assertgo/pkg/assertionsuit"
	agtest "github.com/sku0x20/assertgo/test"
)

func Test_AssertionSuit_Assert_returnsAssertion(t *testing.T) {
	_, s := agtest.NewSink()
	suit := assertionsuit.New(s)
	var result any = suit.Assert("hello")
	if _, ok := result.(*assertion.Assertion[any]); !ok {
		t.Fatal("expected *assertion.Assertion[any]")
	}
}
