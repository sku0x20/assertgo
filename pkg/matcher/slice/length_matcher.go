package slice

import "fmt"

type LengthMatcher[E any] struct {
	length int
}

func NewLengthMatcher[E any](length int) *LengthMatcher[E] {
	return &LengthMatcher[E]{length: length}
}

func (l *LengthMatcher[E]) Match(value []E) bool {
	return len(value) == l.length
}

func (l *LengthMatcher[E]) FailureMsg(value []E) string {
	return fmt.Sprintf("expected length %d but got %d", l.length, len(value))
}
