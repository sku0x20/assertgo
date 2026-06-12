package slice

import (
	"fmt"

	"github.com/sku0x20/assertgo/pkg/comparator"
)

type ContainsMatcher[E any] struct {
	element    E
	comparator comparator.Comparator[E]
}

func NewContainsMatcher[E any](element E) *ContainsMatcher[E] {
	return &ContainsMatcher[E]{
		element:    element,
		comparator: comparator.NewDeepReflectComparator[E](),
	}
}

func (c *ContainsMatcher[E]) Match(value []E) bool {
	for _, v := range value {
		if c.comparator.Compare(v, c.element) == 0 {
			return true
		}
	}
	return false
}

func (c *ContainsMatcher[E]) FailureMsg(value []E) string {
	return fmt.Sprintf("expected %v to contain %v", value, c.element)
}
