package slice

import "fmt"

type EmptyMatcher[E any] struct{}

func NewEmptyMatcher[E any]() *EmptyMatcher[E] {
	return &EmptyMatcher[E]{}
}

func (e *EmptyMatcher[E]) Match(value []E) bool {
	return len(value) == 0
}

func (e *EmptyMatcher[E]) FailureMsg(value []E) string {
	return fmt.Sprintf("expected empty slice but got %v", value)
}
