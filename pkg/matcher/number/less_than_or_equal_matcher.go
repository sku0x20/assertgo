package number

import (
	"fmt"

	"github.com/sku0x20/assertgo/pkg/comparator"
	"github.com/sku0x20/assertgo/pkg/constraints"
)

type LessThanOrEqualMatcher[V constraints.Number] struct {
	other      V
	comparator comparator.Comparator[V]
}

func NewLessThanOrEqualMatcher[V constraints.Number](other V, c comparator.Comparator[V]) *LessThanOrEqualMatcher[V] {
	return &LessThanOrEqualMatcher[V]{other: other, comparator: c}
}

func (l *LessThanOrEqualMatcher[V]) Match(value V) bool {
	return l.comparator.Compare(value, l.other) <= 0
}

func (l *LessThanOrEqualMatcher[V]) FailureMsg(value V) string {
	return fmt.Sprintf("expected %v to be less than or equal to %v", value, l.other)
}
