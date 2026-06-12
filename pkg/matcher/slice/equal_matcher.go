package slice

import (
	"fmt"

	"github.com/sku0x20/assertgo/pkg/comparator"
)

type EqualMatcher[E any] struct {
	other      []E
	comparator comparator.Comparator[E]
}

func NewEqualMatcher[E any](other []E, c comparator.Comparator[E]) *EqualMatcher[E] {
	return &EqualMatcher[E]{other: other, comparator: c}
}

func (e *EqualMatcher[E]) Match(value []E) bool {
	if len(value) != len(e.other) {
		return false
	}
	for i := range value {
		if e.comparator.Compare(value[i], e.other[i]) != 0 {
			return false
		}
	}
	return true
}

func (e *EqualMatcher[E]) FailureMsg(value []E) string {
	return fmt.Sprintf("expected %v but got %v", e.other, value)
}
