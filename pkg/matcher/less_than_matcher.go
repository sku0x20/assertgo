package matcher

import (
	"fmt"

	"github.com/sku0x20/assertgo/pkg/comparator"
)

type LessThanMatcher[V any] struct {
	other      V
	comparator comparator.Comparator[V]
}

func NewLessThanMatcher[V any](other V, c comparator.Comparator[V]) *LessThanMatcher[V] {
	return &LessThanMatcher[V]{other: other, comparator: c}
}

func (l *LessThanMatcher[V]) Match(value V) bool {
	return l.comparator.Compare(value, l.other) < 0
}

func (l *LessThanMatcher[V]) FailureMsg(value V) string {
	return fmt.Sprintf("expected %v to be less than %v", value, l.other)
}
