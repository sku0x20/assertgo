package comparator

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/comparator"
)

func Test_NumberComparator_Compare_Equal(t *testing.T) {
	c := comparator.NewNumberComparator[int]()
	if c.Compare(10, 10) != 0 {
		t.Fatal("expected 0 for equal values")
	}
}

func Test_NumberComparator_Compare_Less(t *testing.T) {
	c := comparator.NewNumberComparator[int]()
	if c.Compare(5, 10) != -1 {
		t.Fatal("expected -1 when a < b")
	}
}

func Test_NumberComparator_Compare_Greater(t *testing.T) {
	c := comparator.NewNumberComparator[int]()
	if c.Compare(10, 5) != 1 {
		t.Fatal("expected 1 when a > b")
	}
}
