package map_matcher

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/comparator"
	map_matcher "github.com/sku0x20/assertgo/pkg/matcher/map"
)

func Test_ContainsAllEntriesMatcher_Match_Pass(t *testing.T) {
	m := map_matcher.NewContainsAllEntriesMatcher(map[string]int{"a": 1}, comparator.NewComparableComparator[int]())
	if !m.Match(map[string]int{"a": 1, "b": 2}) {
		t.Fatal("expected match to be true")
	}
}

func Test_ContainsAllEntriesMatcher_Match_Fail(t *testing.T) {
	m := map_matcher.NewContainsAllEntriesMatcher(map[string]int{"a": 1, "c": 3}, comparator.NewComparableComparator[int]())
	if m.Match(map[string]int{"a": 1, "b": 2}) {
		t.Fatal("expected match to be false")
	}
}

func Test_ContainsAllEntriesMatcher_FailureMsg(t *testing.T) {
	m := map_matcher.NewContainsAllEntriesMatcher(map[string]int{"a": 1}, comparator.NewComparableComparator[int]())
	if m.FailureMsg(map[string]int{"b": 2}) == "" {
		t.Fatal("expected non-empty failure message")
	}
}
