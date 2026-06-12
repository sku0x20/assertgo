package comparator

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/comparator"
)

type point struct{ x, y int }

func Test_DeepReflectComparator_Compare(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		c := comparator.DeepReflectComparator[point]{}
		if c.Compare(point{1, 2}, point{1, 2}) != 0 {
			t.Fatal("expected 0 for equal structs")
		}
	})
	t.Run("fail", func(t *testing.T) {
		c := comparator.DeepReflectComparator[point]{}
		if c.Compare(point{1, 2}, point{3, 4}) == 0 {
			t.Fatal("expected non-zero for unequal structs")
		}
	})
}
