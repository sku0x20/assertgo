package matcher

import "fmt"

type EqualMatcher[V any] struct {
	other      V
	comparator func(a, b V) bool
}

func NewEqualMatcher[V any](other V, comparator func(a, b V) bool) *EqualMatcher[V] {
	return &EqualMatcher[V]{other: other, comparator: comparator}
}

func (e *EqualMatcher[V]) Match(value V) bool {
	return e.comparator(value, e.other)
}

func (e *EqualMatcher[V]) FailureMsg(value V) string {
	return fmt.Sprintf("expected %v but got %v", e.other, value)
}
