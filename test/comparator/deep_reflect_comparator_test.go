package comparator

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/comparator"
)

type point struct{ x, y int }

func Test_DeepReflectComparator_Compare_Pass(t *testing.T) {
	c := comparator.NewDeepReflectComparator[point]()
	if c.Compare(point{1, 2}, point{1, 2}) != 0 {
		t.Fatal("expected 0 for equal structs")
	}
}

func Test_DeepReflectComparator_Compare_Fail(t *testing.T) {
	c := comparator.NewDeepReflectComparator[point]()
	if c.Compare(point{1, 2}, point{3, 4}) == 0 {
		t.Fatal("expected non-zero for unequal structs")
	}
}
