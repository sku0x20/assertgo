package ordered

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/comparator"
	"github.com/sku0x20/assertgo/pkg/matcher/ordered"
)

func Test_GreaterThanMatcher_Match_Pass(t *testing.T) {
	m := ordered.NewGreaterThanMatcher(5, comparator.NewNumberComparator[int]())
	if !m.Match(10) {
		t.Fatal("expected match to be true")
	}
}

func Test_GreaterThanMatcher_Match_Fail(t *testing.T) {
	m := ordered.NewGreaterThanMatcher(5, comparator.NewNumberComparator[int]())
	if m.Match(3) {
		t.Fatal("expected match to be false")
	}
}
