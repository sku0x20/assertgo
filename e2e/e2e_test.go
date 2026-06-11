package e2e

import "testing"
import . "github.com/sku0x20/assertgo"

type mockT struct {
	testing.TB
	failCalled    bool
	failNowCalled bool
}

func (m *mockT) Fail()    { m.failCalled = true }
func (m *mockT) FailNow() { m.failNowCalled = true }

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
