package collection

import (
	"fmt"
	"reflect"
)

type LengthMatcher[T any] struct {
	length int
}

func NewLengthMatcher[T any](length int) *LengthMatcher[T] {
	return &LengthMatcher[T]{length: length}
}

func (l *LengthMatcher[T]) Match(value T) bool {
	return reflect.ValueOf(value).Len() == l.length
}

func (l *LengthMatcher[T]) FailureMsg(value T) string {
	return fmt.Sprintf("expected length %d but got %d", l.length, reflect.ValueOf(value).Len())
}
