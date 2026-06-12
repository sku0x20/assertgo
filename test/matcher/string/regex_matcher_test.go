package string_matcher

import (
	"testing"

	string_matcher "github.com/sku0x20/assertgo/pkg/matcher/string"
)

func Test_RegexMatcher_Match_Pass(t *testing.T) {
	m := string_matcher.NewRegexMatcher(`^\d+$`)
	if !m.Match("12345") {
		t.Fatal("expected match to be true")
	}
}

func Test_RegexMatcher_Match_Fail(t *testing.T) {
	m := string_matcher.NewRegexMatcher(`^\d+$`)
	if m.Match("hello") {
		t.Fatal("expected match to be false")
	}
}

func Test_RegexMatcher_FailureMsg(t *testing.T) {
	m := string_matcher.NewRegexMatcher(`^\d+$`)
	if m.FailureMsg("hello") == "" {
		t.Fatal("expected non-empty failure message")
	}
}
