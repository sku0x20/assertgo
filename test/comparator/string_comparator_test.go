package comparator

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/comparator"
)

func Test_StringComparator_Compare_Equal(t *testing.T) {
	c := comparator.NewStringComparator()
	if c.Compare("hello", "hello") != 0 {
		t.Fatal("expected 0 for equal strings")
	}
}

func Test_StringComparator_Compare_Less(t *testing.T) {
	c := comparator.NewStringComparator()
	if c.Compare("a", "b") != -1 {
		t.Fatal("expected -1 when a < b")
	}
}

func Test_StringComparator_Compare_Greater(t *testing.T) {
	c := comparator.NewStringComparator()
	if c.Compare("b", "a") != 1 {
		t.Fatal("expected 1 when a > b")
	}
}
