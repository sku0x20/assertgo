package number

import (
	"fmt"

	"github.com/sku0x20/assertgo/pkg/comparator"
	"github.com/sku0x20/assertgo/pkg/constraints"
)

type GreaterThanMatcher[V constraints.Number] struct {
	other      V
	comparator comparator.Comparator[V]
}

func NewGreaterThanMatcher[V constraints.Number](other V, c comparator.Comparator[V]) *GreaterThanMatcher[V] {
	return &GreaterThanMatcher[V]{other: other, comparator: c}
}

func (g *GreaterThanMatcher[V]) Match(value V) bool {
	return g.comparator.Compare(value, g.other) > 0
}

func (g *GreaterThanMatcher[V]) FailureMsg(value V) string {
	return fmt.Sprintf("expected %v to be greater than %v", value, g.other)
}
