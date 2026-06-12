package comparator

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/comparator"
)

func Test_ReferenceComparator_Compare_Pass(t *testing.T) {
	p := new(int)
	c := comparator.NewReferenceComparator[*int]()
	if c.Compare(p, p) != 0 {
		t.Fatal("expected 0 for same reference")
	}
}

func Test_ReferenceComparator_Compare_Fail(t *testing.T) {
	c := comparator.NewReferenceComparator[*int]()
	if c.Compare(new(int), new(int)) == 0 {
		t.Fatal("expected non-zero for different references")
	}
}
