package matcher

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/matcher"
)

func Test_BoolMatcher_Match_Pass(t *testing.T) {
	m := matcher.NewBoolMatcher(true)
	if !m.Match(true) {
		t.Fatal("expected match to be true")
	}
}

func Test_BoolMatcher_Match_Fail(t *testing.T) {
	m := matcher.NewBoolMatcher(true)
	if m.Match(false) {
		t.Fatal("expected match to be false")
	}
}

func Test_BoolMatcher_FailureMsg(t *testing.T) {
	m := matcher.NewBoolMatcher(true)
	if m.FailureMsg(false) == "" {
		t.Fatal("expected non-empty failure message")
	}
}
