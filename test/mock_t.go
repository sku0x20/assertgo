package test

import "testing"

type MockT struct {
	testing.TB
	FailCalled    bool
	FailNowCalled bool
	FatalCalled   bool
	FatalMsg      string
}

func (m *MockT) Helper()          {}
func (m *MockT) Fail()            { m.FailCalled = true }
func (m *MockT) FailNow()         { m.FailNowCalled = true }
func (m *MockT) Fatal(args ...any) {
	m.FatalCalled = true
	if len(args) > 0 {
		m.FatalMsg = args[0].(string)
	}
}
