package matcher

import "fmt"

type ReferenceMatcher[V any] struct {
	other V
}

func NewReferenceMatcher[V any](other V) *ReferenceMatcher[V] {
	return &ReferenceMatcher[V]{other: other}
}

func (r *ReferenceMatcher[V]) Match(value V) bool {
	return any(value) == any(r.other)
}

func (r *ReferenceMatcher[V]) FailureMsg(value V) string {
	return fmt.Sprintf("expected same reference as %v but got %v", r.other, value)
}
