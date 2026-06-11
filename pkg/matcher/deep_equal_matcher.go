package matcher

import (
	"fmt"
	"reflect"
)

type DeepEqualMatcher struct {
	other any
}

func NewDeepEqualMatcher(other any) *DeepEqualMatcher {
	return &DeepEqualMatcher{other: other}
}

func (d *DeepEqualMatcher) Match(value any) bool {
	return reflect.DeepEqual(value, d.other)
}

func (d *DeepEqualMatcher) FailureMsg(value any) string {
	return fmt.Sprintf("expected %v but got %v", d.other, value)
}
