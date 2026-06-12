package collection

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/matcher/collection"
)

func Test_LengthMatcher_Match_Pass_Slice(t *testing.T) {
	m := collection.NewLengthMatcher[[]int](3)
	if !m.Match([]int{1, 2, 3}) {
		t.Fatal("expected match to be true")
	}
}

func Test_LengthMatcher_Match_Pass_Map(t *testing.T) {
	m := collection.NewLengthMatcher[map[string]int](2)
	if !m.Match(map[string]int{"a": 1, "b": 2}) {
		t.Fatal("expected match to be true")
	}
}

func Test_LengthMatcher_Match_Pass_String(t *testing.T) {
	m := collection.NewLengthMatcher[string](5)
	if !m.Match("hello") {
		t.Fatal("expected match to be true")
	}
}

func Test_LengthMatcher_Match_Fail(t *testing.T) {
	m := collection.NewLengthMatcher[[]int](3)
	if m.Match([]int{1, 2}) {
		t.Fatal("expected match to be false")
	}
}

func Test_LengthMatcher_FailureMsg(t *testing.T) {
	m := collection.NewLengthMatcher[[]int](3)
	if m.FailureMsg([]int{1, 2}) == "" {
		t.Fatal("expected non-empty failure message")
	}
}
