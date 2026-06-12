package map_matcher

import (
	"fmt"

	"github.com/sku0x20/assertgo/pkg/comparator"
)

type ContainsAllEntriesMatcher[K comparable, V any] struct {
	entries    map[K]V
	comparator comparator.Comparator[V]
}

func NewContainsAllEntriesMatcher[K comparable, V any](entries map[K]V, c comparator.Comparator[V]) *ContainsAllEntriesMatcher[K, V] {
	return &ContainsAllEntriesMatcher[K, V]{entries: entries, comparator: c}
}

func (c *ContainsAllEntriesMatcher[K, V]) Match(value map[K]V) bool {
	for k, v := range c.entries {
		val, ok := value[k]
		if !ok || c.comparator.Compare(val, v) != 0 {
			return false
		}
	}
	return true
}

func (c *ContainsAllEntriesMatcher[K, V]) FailureMsg(value map[K]V) string {
	return fmt.Sprintf("expected %v to contain all entries %v", value, c.entries)
}
