package comparator

import "github.com/sku0x20/assertgo/pkg/constraints"

type NumberComparator[V constraints.Number] struct{}

func NewNumberComparator[V constraints.Number]() *NumberComparator[V] {
	return &NumberComparator[V]{}
}

func (c *NumberComparator[V]) Compare(a, b V) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}
