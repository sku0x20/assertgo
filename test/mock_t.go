package test

import "testing"

type mockT struct {
	testing.TB
	fatalCalled bool
	fatalMsg    string
}

func (m *mockT) Helper() {}

func (m *mockT) Fatal(args ...any) {
	m.fatalCalled = true
	if len(args) > 0 {
		m.fatalMsg = args[0].(string)
	}
}
