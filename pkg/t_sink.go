package pkg

import "testing"

type TSink struct {
	t testing.TB
}

func NewTSink(t testing.TB) *TSink {
	return &TSink{t: t}
}

func (f *TSink) Fail(msg string) {
	f.t.Helper()
	f.t.Fatal(msg)
}
