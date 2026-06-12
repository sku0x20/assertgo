package slice

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/matcher/slice"
)

func Test_LengthMatcher_Match_Pass(t *testing.T) {
	m := slice.NewLengthMatcher[int](3)
	if !m.Match([]int{1, 2, 3}) {
		t.Fatal("expected match to be true")
	}
}

func Test_LengthMatcher_Match_Fail(t *testing.T) {
	m := slice.NewLengthMatcher[int](3)
	if m.Match([]int{1, 2}) {
		t.Fatal("expected match to be false")
	}
}

func Test_LengthMatcher_FailureMsg(t *testing.T) {
	m := slice.NewLengthMatcher[int](3)
	if m.FailureMsg([]int{1, 2}) == "" {
		t.Fatal("expected non-empty failure message")
	}
}
