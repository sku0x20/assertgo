package matcher

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/matcher"
)

func Test_NilMatcher_Match(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		m := matcher.NewNilMatcher()
		if !m.Match(nil) {
			t.Fatal("expected match to be true")
		}
	})
	t.Run("fail", func(t *testing.T) {
		m := matcher.NewNilMatcher()
		if m.Match("hello") {
			t.Fatal("expected match to be false")
		}
	})
}

func Test_NilMatcher_FailureMsg(t *testing.T) {
	m := matcher.NewNilMatcher()
	if m.FailureMsg("hello") == "" {
		t.Fatal("expected non-empty failure message")
	}
}
