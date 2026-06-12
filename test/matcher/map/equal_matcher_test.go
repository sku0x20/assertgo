package map_matcher

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/comparator"
	map_matcher "github.com/sku0x20/assertgo/pkg/matcher/map"
)

func Test_MapEqualMatcher_Match_Pass(t *testing.T) {
	m := map_matcher.NewEqualMatcher(map[string]int{"a": 1, "b": 2}, comparator.NewComparableComparator[int]())
	if !m.Match(map[string]int{"a": 1, "b": 2}) {
		t.Fatal("expected match to be true")
	}
}

func Test_MapEqualMatcher_Match_Fail_DifferentValues(t *testing.T) {
	m := map_matcher.NewEqualMatcher(map[string]int{"a": 1, "b": 2}, comparator.NewComparableComparator[int]())
	if m.Match(map[string]int{"a": 1, "b": 3}) {
		t.Fatal("expected match to be false")
	}
}

func Test_MapEqualMatcher_Match_Fail_DifferentKeys(t *testing.T) {
	m := map_matcher.NewEqualMatcher(map[string]int{"a": 1, "b": 2}, comparator.NewComparableComparator[int]())
	if m.Match(map[string]int{"a": 1, "c": 2}) {
		t.Fatal("expected match to be false")
	}
}

func Test_MapEqualMatcher_FailureMsg(t *testing.T) {
	m := map_matcher.NewEqualMatcher(map[string]int{"a": 1}, comparator.NewComparableComparator[int]())
	if m.FailureMsg(map[string]int{"a": 2}) == "" {
		t.Fatal("expected non-empty failure message")
	}
}
