package pkg

import "testing"

type FailSink struct {
	t    testing.TB
	hard bool
}

func (f *FailSink) Fail(msg string) {
	f.t.Helper()
	if f.hard {
		f.t.Fatal(msg)
	} else {
		f.t.Error(msg)
	}
}
