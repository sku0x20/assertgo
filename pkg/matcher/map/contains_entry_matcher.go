package map_matcher

import (
	"fmt"

	"github.com/sku0x20/assertgo/pkg/comparator"
)

type ContainsEntryMatcher[K comparable, V any] struct {
	key        K
	value      V
	comparator comparator.Comparator[V]
}

func NewContainsEntryMatcher[K comparable, V any](key K, value V, c comparator.Comparator[V]) *ContainsEntryMatcher[K, V] {
	return &ContainsEntryMatcher[K, V]{key: key, value: value, comparator: c}
}

func (c *ContainsEntryMatcher[K, V]) Match(value map[K]V) bool {
	v, ok := value[c.key]
	return ok && c.comparator.Compare(v, c.value) == 0
}

func (c *ContainsEntryMatcher[K, V]) FailureMsg(value map[K]V) string {
	return fmt.Sprintf("expected %v to contain entry {%v: %v}", value, c.key, c.value)
}
