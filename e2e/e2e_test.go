package e2e

import (
	"testing"

	. "github.com/sku0x20/assertgo"
	agtest "github.com/sku0x20/assertgo/test"
)

func Test_FailHard(t *testing.T) {
	mock := &agtest.MockT{}
	T(mock)
	if !mock.FatalCalled {
		t.Fatal("should have called FailNow")
	}
}
