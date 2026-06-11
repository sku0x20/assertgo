package e2e

import (
	"testing"

	. "github.com/sku0x20/assertgo"
	agtest "github.com/sku0x20/assertgo/test"
)

func Test_Assert(tt *testing.T) {
	t := &agtest.MockT{}
	T(t).Assert("10").EqualTo("9")
	if !t.FatalCalled {
		tt.Fatal("should have called FailNow")
	}
}
