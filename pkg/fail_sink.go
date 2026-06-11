package pkg

import "testing"

type FailSink struct {
	t testing.TB
}

func (f *FailSink) Fail(msg string) {
	f.t.Helper()
	f.t.Fatal(msg)
}
