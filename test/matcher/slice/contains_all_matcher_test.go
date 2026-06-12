package slice

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/comparator"
	"github.com/sku0x20/assertgo/pkg/matcher/slice"
)

func Test_ContainsAllMatcher_Match_Pass(t *testing.T) {
	m := slice.NewContainsAllMatcher([]int{1, 3}, comparator.NewComparableComparator[int]())
	if !m.Match([]int{1, 2, 3}) {
		t.Fatal("expected match to be true")
	}
}

func Test_ContainsAllMatcher_Match_Fail(t *testing.T) {
	m := slice.NewContainsAllMatcher([]int{1, 4}, comparator.NewComparableComparator[int]())
	if m.Match([]int{1, 2, 3}) {
		t.Fatal("expected match to be false")
	}
}

func Test_ContainsAllMatcher_FailureMsg(t *testing.T) {
	m := slice.NewContainsAllMatcher([]int{1, 4}, comparator.NewComparableComparator[int]())
	if m.FailureMsg([]int{1, 2, 3}) == "" {
		t.Fatal("expected non-empty failure message")
	}
}
