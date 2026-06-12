package integer

import (
	"fmt"

	"github.com/sku0x20/assertgo/pkg/comparator"
	"github.com/sku0x20/assertgo/pkg/constraints"
)

type InRangeMatcher[V constraints.Integer] struct {
	min, max   V
	comparator comparator.Comparator[V]
}

func NewInRangeMatcher[V constraints.Integer](min, max V, c comparator.Comparator[V]) *InRangeMatcher[V] {
	return &InRangeMatcher[V]{min: min, max: max, comparator: c}
}

func (r *InRangeMatcher[V]) Match(value V) bool {
	return r.comparator.Compare(value, r.min) >= 0 && r.comparator.Compare(value, r.max) <= 0
}

func (r *InRangeMatcher[V]) FailureMsg(value V) string {
	return fmt.Sprintf("expected %v to be in range [%v, %v]", value, r.min, r.max)
}
