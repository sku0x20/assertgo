package slice

import (
	"fmt"

	"github.com/sku0x20/assertgo/pkg/comparator"
)

type UniqueMatcher[E any] struct {
	comparator comparator.Comparator[E]
}

func NewUniqueMatcher[E any](c comparator.Comparator[E]) *UniqueMatcher[E] {
	return &UniqueMatcher[E]{comparator: c}
}

func (u *UniqueMatcher[E]) Match(value []E) bool {
	for i := 0; i < len(value); i++ {
		for j := i + 1; j < len(value); j++ {
			if u.comparator.Compare(value[i], value[j]) == 0 {
				return false
			}
		}
	}
	return true
}

func (u *UniqueMatcher[E]) FailureMsg(value []E) string {
	return fmt.Sprintf("expected all elements to be unique but got %v", value)
}
