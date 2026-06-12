package map_matcher

import (
	"testing"

	map_matcher "github.com/sku0x20/assertgo/pkg/matcher/map"
)

func Test_MapLengthMatcher_Match_Pass(t *testing.T) {
	m := map_matcher.NewLengthMatcher[string, int](2)
	if !m.Match(map[string]int{"a": 1, "b": 2}) {
		t.Fatal("expected match to be true")
	}
}

func Test_MapLengthMatcher_Match_Fail(t *testing.T) {
	m := map_matcher.NewLengthMatcher[string, int](2)
	if m.Match(map[string]int{"a": 1}) {
		t.Fatal("expected match to be false")
	}
}

func Test_MapLengthMatcher_FailureMsg(t *testing.T) {
	m := map_matcher.NewLengthMatcher[string, int](2)
	if m.FailureMsg(map[string]int{"a": 1}) == "" {
		t.Fatal("expected non-empty failure message")
	}
}
