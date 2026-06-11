package test

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg"
	"github.com/sku0x20/assertgo/pkg/assertions"
	"github.com/sku0x20/assertgo/pkg/sink"
)

func Test_AssertionSuite_Assert_returnsAnyAssertions(t *testing.T) {
	mock := &MockT{}
	s := sink.NewTSink(mock)
	suite := pkg.NewAssertionSuite(s)
	var result any = suite.Assert("hello")
	if _, ok := result.(*assertions.AnyAssertions); !ok {
		t.Fatal("expected *assertions.AnyAssertions")
	}
}
