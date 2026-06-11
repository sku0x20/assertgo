package test

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg"
	"github.com/sku0x20/assertgo/pkg/assertions"
)

func Test_AssertionSuite_Assert_returnsAnyAssertions(t *testing.T) {
	_, s := NewSink()
	suite := pkg.NewAssertionSuite(s)
	var result any = suite.Assert("hello")
	if _, ok := result.(*assertions.AnyAssertions); !ok {
		t.Fatal("expected *assertions.AnyAssertions")
	}
}
