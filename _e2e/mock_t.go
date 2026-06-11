package e2e

import "testing"

type mockT struct {
	testing.TB
	failCalled    bool
	failNowCalled bool
}

func (m *mockT) Fail()    { m.failCalled = true }
func (m *mockT) FailNow() { m.failNowCalled = true }
