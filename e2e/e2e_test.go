package e2e

import (
	"testing"

	. "github.com/sku0x20/assertgo"
	agtest "github.com/sku0x20/assertgo/test"
)

func Test_FailHard(t *testing.T) {
	mock := &agtest.MockT{}
	T(mock)
	if !mock.FailNowCalled {
		t.Fatal("should have called FailNow")
	}
}

func Test_FailSoft(t *testing.T) {
	mock := &agtest.MockT{}
	Ts(mock)
	if !mock.FailCalled {
		t.Fatal("should have called Fail")
	}
}
