package e2e

import "testing"
import . "github.com/sku0x20/assertgo"

type mockT struct {
	testing.TB
	failed bool
}

func (m *mockT) Fail() { m.failed = true }

func Test_FailHard(t *testing.T) {
	T(t)
}

func Test_FailSoft(t *testing.T) {
	mock := &mockT{}
	Ts(mock)
	if !mock.failed {
		t.Fatal("should have failed")
	}
}
