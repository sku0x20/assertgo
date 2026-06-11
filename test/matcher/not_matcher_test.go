package matcher

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/matcher"
)

func Test_NotMatcher_Match(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		m := matcher.NewNotMatcher[any](&mockMatcher{matchResult: false})
		if !m.Match(nil) {
			t.Fatal("expected match to be true")
		}
	})
	t.Run("fail", func(t *testing.T) {
		m := matcher.NewNotMatcher[any](&mockMatcher{matchResult: true})
		if m.Match(nil) {
			t.Fatal("expected match to be false")
		}
	})
}

func Test_NotMatcher_FailureMsg(t *testing.T) {
	m := matcher.NewNotMatcher[any](&mockMatcher{failureMsg: "some failure"})
	if m.FailureMsg(nil) != "not: some failure" {
		t.Fatal("expected failure message to be prefixed with 'not: '")
	}
}
