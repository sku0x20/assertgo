package matcher

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/matcher"
)

func Test_TypeMatcher_Match_Pass(t *testing.T) {
	m := matcher.NewTypeMatcher[int]()
	if !m.Match(42) {
		t.Fatal("expected match to be true")
	}
}

func Test_TypeMatcher_Match_Fail(t *testing.T) {
	m := matcher.NewTypeMatcher[int]()
	if m.Match("hello") {
		t.Fatal("expected match to be false")
	}
}

func Test_TypeMatcher_FailureMsg(t *testing.T) {
	m := matcher.NewTypeMatcher[int]()
	if m.FailureMsg("hello") == "" {
		t.Fatal("expected non-empty failure message")
	}
}
