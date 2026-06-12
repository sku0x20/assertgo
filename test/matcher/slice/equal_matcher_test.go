package slice

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/comparator"
	"github.com/sku0x20/assertgo/pkg/matcher/slice"
)

func Test_SliceEqualMatcher_Match_Pass(t *testing.T) {
	m := slice.NewEqualMatcher([]int{1, 2, 3}, comparator.NewComparableComparator[int]())
	if !m.Match([]int{1, 2, 3}) {
		t.Fatal("expected match to be true")
	}
}

func Test_SliceEqualMatcher_Match_Fail_DifferentElements(t *testing.T) {
	m := slice.NewEqualMatcher([]int{1, 2, 3}, comparator.NewComparableComparator[int]())
	if m.Match([]int{1, 2, 4}) {
		t.Fatal("expected match to be false")
	}
}

func Test_SliceEqualMatcher_Match_Fail_DifferentLength(t *testing.T) {
	m := slice.NewEqualMatcher([]int{1, 2, 3}, comparator.NewComparableComparator[int]())
	if m.Match([]int{1, 2}) {
		t.Fatal("expected match to be false")
	}
}

func Test_SliceEqualMatcher_FailureMsg(t *testing.T) {
	m := slice.NewEqualMatcher([]int{1, 2, 3}, comparator.NewComparableComparator[int]())
	if m.FailureMsg([]int{1, 2, 4}) == "" {
		t.Fatal("expected non-empty failure message")
	}
}
