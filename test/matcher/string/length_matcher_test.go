package string_matcher

import (
	"testing"

	string_matcher "github.com/sku0x20/assertgo/pkg/matcher/string"
)

func Test_StringLengthMatcher_Match_Pass(t *testing.T) {
	m := string_matcher.NewLengthMatcher(5)
	if !m.Match("hello") {
		t.Fatal("expected match to be true")
	}
}

func Test_StringLengthMatcher_Match_Fail(t *testing.T) {
	m := string_matcher.NewLengthMatcher(5)
	if m.Match("hi") {
		t.Fatal("expected match to be false")
	}
}

func Test_StringLengthMatcher_FailureMsg(t *testing.T) {
	m := string_matcher.NewLengthMatcher(5)
	if m.FailureMsg("hi") == "" {
		t.Fatal("expected non-empty failure message")
	}
}
