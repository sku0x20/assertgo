package slice

import (
	"fmt"

	"github.com/sku0x20/assertgo/pkg/comparator"
)

type ContainsNoneMatcher[E any] struct {
	elements   []E
	comparator comparator.Comparator[E]
}

func NewContainsNoneMatcher[E any](elements []E, c comparator.Comparator[E]) *ContainsNoneMatcher[E] {
	return &ContainsNoneMatcher[E]{elements: elements, comparator: c}
}

func (c *ContainsNoneMatcher[E]) Match(value []E) bool {
	for _, element := range c.elements {
		for _, v := range value {
			if c.comparator.Compare(v, element) == 0 {
				return false
			}
		}
	}
	return true
}

func (c *ContainsNoneMatcher[E]) FailureMsg(value []E) string {
	return fmt.Sprintf("expected %v to contain none of %v", value, c.elements)
}
