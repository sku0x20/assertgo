package matcher

import (
	"fmt"

	"github.com/sku0x20/assertgo/pkg/comparator"
)

type GreaterThanOrEqualMatcher[V any] struct {
	other      V
	comparator comparator.Comparator[V]
}

func NewGreaterThanOrEqualMatcher[V any](other V, c comparator.Comparator[V]) *GreaterThanOrEqualMatcher[V] {
	return &GreaterThanOrEqualMatcher[V]{other: other, comparator: c}
}

func (g *GreaterThanOrEqualMatcher[V]) Match(value V) bool {
	return g.comparator.Compare(value, g.other) >= 0
}

func (g *GreaterThanOrEqualMatcher[V]) FailureMsg(value V) string {
	return fmt.Sprintf("expected %v to be greater than or equal to %v", value, g.other)
}
