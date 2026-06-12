package slice

import (
	"fmt"

	"github.com/sku0x20/assertgo/pkg/comparator"
)

type ContainsAllMatcher[E any] struct {
	elements   []E
	comparator comparator.Comparator[E]
}

func NewContainsAllMatcher[E any](elements []E, c comparator.Comparator[E]) *ContainsAllMatcher[E] {
	return &ContainsAllMatcher[E]{elements: elements, comparator: c}
}

func (c *ContainsAllMatcher[E]) Match(value []E) bool {
	for _, element := range c.elements {
		found := false
		for _, v := range value {
			if c.comparator.Compare(v, element) == 0 {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (c *ContainsAllMatcher[E]) FailureMsg(value []E) string {
	return fmt.Sprintf("expected %v to contain all of %v", value, c.elements)
}
