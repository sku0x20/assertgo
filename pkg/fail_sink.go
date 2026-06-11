package pkg

import "testing"

type FailSink struct {
	t testing.TB
}

func NewFailSink(t testing.TB) *FailSink {
	return &FailSink{t: t}
}

func (f *FailSink) Fail(msg string) {
	f.t.Helper()
	f.t.Fatal(msg)
}
