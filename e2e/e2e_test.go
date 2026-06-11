package e2e

import (
	"testing"

	. "github.com/sku0x20/assertgo"
)

func Test_FailHard(t *testing.T) {
	mock := &mockT{}
	T(mock)
	if !mock.failNowCalled {
		t.Fatal("should have called FailNow")
	}
}

func Test_FailSoft(t *testing.T) {
	mock := &mockT{}
	Ts(mock)
	if !mock.failCalled {
		t.Fatal("should have called Fail")
	}
}
