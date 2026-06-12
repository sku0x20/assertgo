package matcher

import (
	"reflect"
	"testing"

	"github.com/sku0x20/assertgo/pkg/matcher"
)

func Test_EqualMatcher_Match(t *testing.T) {
	t.Run("pass with custom comparator", func(t *testing.T) {
		m := matcher.NewEqualMatcher("hello", func(a, b string) bool { return a == b })
		if !m.Match("hello") {
			t.Fatal("expected match to be true")
		}
	})
	t.Run("fail with custom comparator", func(t *testing.T) {
		m := matcher.NewEqualMatcher("hello", func(a, b string) bool { return a == b })
		if m.Match("world") {
			t.Fatal("expected match to be false")
		}
	})
	t.Run("pass with DeepEqual comparator", func(t *testing.T) {
		m := matcher.NewEqualMatcher([]int{1, 2, 3}, reflect.DeepEqual)
		if !m.Match([]int{1, 2, 3}) {
			t.Fatal("expected match to be true")
		}
	})
	t.Run("fail with DeepEqual comparator", func(t *testing.T) {
		m := matcher.NewEqualMatcher([]int{1, 2, 3}, reflect.DeepEqual)
		if m.Match([]int{1, 2, 4}) {
			t.Fatal("expected match to be false")
		}
	})
}

func Test_EqualMatcher_FailureMsg(t *testing.T) {
	m := matcher.NewEqualMatcher("hello", func(a, b string) bool { return a == b })
	if m.FailureMsg("world") == "" {
		t.Fatal("expected non-empty failure message")
	}
}
