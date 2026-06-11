package test

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg"
	"github.com/sku0x20/assertgo/pkg/assertion"
)

func Test_AssertionSuite_Assert_returnsAnyAssertion(t *testing.T) {
	_, s := NewSink()
	suite := pkg.NewAssertionSuite(s)
	var result any = suite.Assert("hello")
	if _, ok := result.(*assertion.AnyAssertion); !ok {
		t.Fatal("expected *assertion.AnyAssertion")
	}
}
