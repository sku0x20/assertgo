package comparator

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/comparator"
)

func Test_DeepReflectComparator_Compare(t *testing.T) {
	t.Run("equal strings", func(t *testing.T) {
		c := comparator.DeepReflectComparator[string]{}
		if c.Compare("hello", "hello") != 0 {
			t.Fatal("expected 0 for equal values")
		}
	})
	t.Run("unequal strings", func(t *testing.T) {
		c := comparator.DeepReflectComparator[string]{}
		if c.Compare("hello", "world") == 0 {
			t.Fatal("expected non-zero for unequal values")
		}
	})
	t.Run("equal slices", func(t *testing.T) {
		c := comparator.DeepReflectComparator[[]int]{}
		if c.Compare([]int{1, 2, 3}, []int{1, 2, 3}) != 0 {
			t.Fatal("expected 0 for equal slices")
		}
	})
	t.Run("unequal slices", func(t *testing.T) {
		c := comparator.DeepReflectComparator[[]int]{}
		if c.Compare([]int{1, 2, 3}, []int{1, 2, 4}) == 0 {
			t.Fatal("expected non-zero for unequal slices")
		}
	})
	t.Run("equal structs", func(t *testing.T) {
		type point struct{ x, y int }
		c := comparator.DeepReflectComparator[point]{}
		if c.Compare(point{1, 2}, point{1, 2}) != 0 {
			t.Fatal("expected 0 for equal structs")
		}
	})
	t.Run("unequal structs", func(t *testing.T) {
		type point struct{ x, y int }
		c := comparator.DeepReflectComparator[point]{}
		if c.Compare(point{1, 2}, point{3, 4}) == 0 {
			t.Fatal("expected non-zero for unequal structs")
		}
	})
}
