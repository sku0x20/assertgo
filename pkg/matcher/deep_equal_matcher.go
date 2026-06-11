package matcher

import (
	"fmt"
	"reflect"
)

type DeepEqualMatcher[V any] struct {
	other V
}

func NewDeepEqualMatcher[V any](other V) *DeepEqualMatcher[V] {
	return &DeepEqualMatcher[V]{other: other}
}

func (d *DeepEqualMatcher[V]) Match(value V) bool {
	return reflect.DeepEqual(value, d.other)
}

func (d *DeepEqualMatcher[V]) FailureMsg(value V) string {
	return fmt.Sprintf("expected %v but got %v", d.other, value)
}
