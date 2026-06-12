package matcher

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/matcher"
)

func Test_EqualMatcher_Match_Pass(t *testing.T) {
	m := matcher.NewEqualMatcher("hello", &MockComparator[string]{Result: 0})
	if !m.Match("hello") {
		t.Fatal("expected match to be true")
	}
}

func Test_EqualMatcher_Match_Fail(t *testing.T) {
	m := matcher.NewEqualMatcher("hello", &MockComparator[string]{Result: -1})
	if m.Match("world") {
		t.Fatal("expected match to be false")
	}
}

func Test_EqualMatcher_FailureMsg(t *testing.T) {
	m := matcher.NewEqualMatcher("hello", &MockComparator[string]{})
	if m.FailureMsg("world") == "" {
		t.Fatal("expected non-empty failure message")
	}
}
