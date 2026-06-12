package comparator

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/comparator"
)

func Test_ComparableComparator_Compare_Pass(t *testing.T) {
	c := comparator.NewComparableComparator[int]()
	if c.Compare(10, 10) != 0 {
		t.Fatal("expected 0 for equal values")
	}
}

func Test_ComparableComparator_Compare_Fail(t *testing.T) {
	c := comparator.NewComparableComparator[int]()
	if c.Compare(10, 20) == 0 {
		t.Fatal("expected non-zero for different values")
	}
}
