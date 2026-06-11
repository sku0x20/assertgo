package matcher

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/matcher"
)

func Test_DeepEqualMatcher_Match(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		m := matcher.NewDeepEqualMatcher("hello")
		if !m.Match("hello") {
			t.Fatal("expected match to be true")
		}
	})
	t.Run("fail", func(t *testing.T) {
		m := matcher.NewDeepEqualMatcher("hello")
		if m.Match("world") {
			t.Fatal("expected match to be false")
		}
	})
}

func Test_DeepEqualMatcher_FailureMsg(t *testing.T) {
	m := matcher.NewDeepEqualMatcher("hello")
	if m.FailureMsg("world") == "" {
		t.Fatal("expected non-empty failure message")
	}
}
