package ordered

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/comparator"
	"github.com/sku0x20/assertgo/pkg/matcher/ordered"
)

func Test_LessThanMatcher_Match_Pass(t *testing.T) {
	m := ordered.NewLessThanMatcher(10, comparator.NewNumberComparator[int]())
	if !m.Match(5) {
		t.Fatal("expected match to be true")
	}
}

func Test_LessThanMatcher_Match_Fail(t *testing.T) {
	m := ordered.NewLessThanMatcher(10, comparator.NewNumberComparator[int]())
	if m.Match(15) {
		t.Fatal("expected match to be false")
	}
}
