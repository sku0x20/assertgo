package e2e

import (
	"testing"

	. "github.com/sku0x20/assertgo"
	agtest "github.com/sku0x20/assertgo/test"
)

func Test_Fail(tt *testing.T) {
	t := &agtest.MockT{}
	//T(t).Assert("10").isEqualTo("9")
	T(t)
	if !t.FatalCalled {
		t.Fatal("should have called FailNow")
	}
}
