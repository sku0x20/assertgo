package matcher

import "fmt"

type TypeMatcher[T any] struct{}

func NewTypeMatcher[T any]() *TypeMatcher[T] {
	return &TypeMatcher[T]{}
}

func (t *TypeMatcher[T]) Match(value any) bool {
	_, ok := value.(T)
	return ok
}

func (t *TypeMatcher[T]) FailureMsg(value any) string {
	return fmt.Sprintf("expected type %T but got %T", *new(T), value)
}
