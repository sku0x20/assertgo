package map_matcher

import "fmt"

type LengthMatcher[K comparable, V any] struct {
	length int
}

func NewLengthMatcher[K comparable, V any](length int) *LengthMatcher[K, V] {
	return &LengthMatcher[K, V]{length: length}
}

func (l *LengthMatcher[K, V]) Match(value map[K]V) bool {
	return len(value) == l.length
}

func (l *LengthMatcher[K, V]) FailureMsg(value map[K]V) string {
	return fmt.Sprintf("expected length %d but got %d", l.length, len(value))
}
