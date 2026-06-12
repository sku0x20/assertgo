package matcher

import (
	"fmt"

	"github.com/sku0x20/assertgo/pkg/comparator"
)

type EqualMatcher[V any] struct {
	other      V
	comparator comparator.Comparator[V]
}

func NewEqualMatcher[V any](other V, c comparator.Comparator[V]) *EqualMatcher[V] {
	return &EqualMatcher[V]{other: other, comparator: c}
}

func (e *EqualMatcher[V]) Match(value V) bool {
	return e.comparator.Compare(value, e.other) == 0
}

func (e *EqualMatcher[V]) FailureMsg(value V) string {
	return fmt.Sprintf("expected %v but got %v", e.other, value)
}
