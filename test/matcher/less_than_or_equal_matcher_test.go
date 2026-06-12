package matcher

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/comparator"
	"github.com/sku0x20/assertgo/pkg/matcher"
)

func Test_LessThanOrEqualMatcher_Match_Pass(t *testing.T) {
	m := matcher.NewLessThanOrEqualMatcher(10, comparator.NewNumberComparator[int]())
	if !m.Match(10) {
		t.Fatal("expected match to be true")
	}
}

func Test_LessThanOrEqualMatcher_Match_Fail(t *testing.T) {
	m := matcher.NewLessThanOrEqualMatcher(10, comparator.NewNumberComparator[int]())
	if m.Match(15) {
		t.Fatal("expected match to be false")
	}
}
