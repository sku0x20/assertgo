package string_matcher

import (
	"testing"

	string_matcher "github.com/sku0x20/assertgo/pkg/matcher/string"
)

func Test_HasPrefixMatcher_Match_Pass(t *testing.T) {
	m := string_matcher.NewHasPrefixMatcher("hel")
	if !m.Match("hello") {
		t.Fatal("expected match to be true")
	}
}

func Test_HasPrefixMatcher_Match_Fail(t *testing.T) {
	m := string_matcher.NewHasPrefixMatcher("xyz")
	if m.Match("hello") {
		t.Fatal("expected match to be false")
	}
}

func Test_HasPrefixMatcher_FailureMsg(t *testing.T) {
	m := string_matcher.NewHasPrefixMatcher("xyz")
	if m.FailureMsg("hello") == "" {
		t.Fatal("expected non-empty failure message")
	}
}
