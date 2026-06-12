package matcher

import (
	"reflect"
	"testing"

	"github.com/sku0x20/assertgo/pkg/matcher"
)

type stringComparator struct{}

func (s stringComparator) Compare(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

type deepEqualComparator[V any] struct{}

func (d deepEqualComparator[V]) Compare(a, b V) int {
	if reflect.DeepEqual(a, b) {
		return 0
	}
	return -1
}

func Test_EqualMatcher_Match(t *testing.T) {
	t.Run("pass with custom comparator", func(t *testing.T) {
		m := matcher.NewEqualMatcher("hello", stringComparator{})
		if !m.Match("hello") {
			t.Fatal("expected match to be true")
		}
	})
	t.Run("fail with custom comparator", func(t *testing.T) {
		m := matcher.NewEqualMatcher("hello", stringComparator{})
		if m.Match("world") {
			t.Fatal("expected match to be false")
		}
	})
	t.Run("pass with DeepEqual comparator", func(t *testing.T) {
		m := matcher.NewEqualMatcher([]int{1, 2, 3}, deepEqualComparator[[]int]{})
		if !m.Match([]int{1, 2, 3}) {
			t.Fatal("expected match to be true")
		}
	})
	t.Run("fail with DeepEqual comparator", func(t *testing.T) {
		m := matcher.NewEqualMatcher([]int{1, 2, 3}, deepEqualComparator[[]int]{})
		if m.Match([]int{1, 2, 4}) {
			t.Fatal("expected match to be false")
		}
	})
}

func Test_EqualMatcher_FailureMsg(t *testing.T) {
	m := matcher.NewEqualMatcher("hello", stringComparator{})
	if m.FailureMsg("world") == "" {
		t.Fatal("expected non-empty failure message")
	}
}
