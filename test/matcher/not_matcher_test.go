package matcher

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/matcher"
)

func Test_NotMatcher_Match(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		m := matcher.NewNotMatcher[any](&MockMatcher{MatchResult: false})
		if !m.Match(nil) {
			t.Fatal("expected match to be true")
		}
	})
	t.Run("fail", func(t *testing.T) {
		m := matcher.NewNotMatcher[any](&MockMatcher{MatchResult: true})
		if m.Match(nil) {
			t.Fatal("expected match to be false")
		}
	})
}

func Test_NotMatcher_FailureMsg(t *testing.T) {
	m := matcher.NewNotMatcher[any](&MockMatcher{Msg: "some failure"})
	msg := m.FailureMsg(nil)
	if msg != "not: some failure" {
		t.Fatal("expected failure message to be prefixed with 'not: '")
	}
}

func Test_NotMatcher_Set(t *testing.T) {
	m := matcher.NewNotMatcher[any](nil)
	m.Set(&MockMatcher{MatchResult: false})
	if !m.Match(nil) {
		t.Fatal("expected match to be true after Set")
	}
}
