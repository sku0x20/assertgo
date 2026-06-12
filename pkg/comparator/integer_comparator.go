package comparator

import "github.com/sku0x20/assertgo/pkg/constraints"

type IntegerComparator[V constraints.Integer] struct{}

func NewIntegerComparator[V constraints.Integer]() *IntegerComparator[V] {
	return &IntegerComparator[V]{}
}

func (c *IntegerComparator[V]) Compare(a, b V) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}
