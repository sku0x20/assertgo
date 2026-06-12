package map_matcher

import (
	"fmt"

	"github.com/sku0x20/assertgo/pkg/comparator"
)

type EqualMatcher[K comparable, V any] struct {
	other      map[K]V
	comparator comparator.Comparator[V]
}

func NewEqualMatcher[K comparable, V any](other map[K]V, c comparator.Comparator[V]) *EqualMatcher[K, V] {
	return &EqualMatcher[K, V]{other: other, comparator: c}
}

func (e *EqualMatcher[K, V]) Match(value map[K]V) bool {
	if len(value) != len(e.other) {
		return false
	}
	for k, v := range e.other {
		val, ok := value[k]
		if !ok || e.comparator.Compare(val, v) != 0 {
			return false
		}
	}
	return true
}

func (e *EqualMatcher[K, V]) FailureMsg(value map[K]V) string {
	return fmt.Sprintf("expected %v but got %v", e.other, value)
}
