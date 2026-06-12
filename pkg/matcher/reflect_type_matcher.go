package matcher

import (
	"fmt"
	"reflect"
)

type ReflectTypeMatcher struct {
	expected reflect.Type
}

func NewReflectTypeMatcher(expected reflect.Type) *ReflectTypeMatcher {
	return &ReflectTypeMatcher{expected: expected}
}

func (r *ReflectTypeMatcher) Match(value any) bool {
	return reflect.TypeOf(value) == r.expected
}

func (r *ReflectTypeMatcher) FailureMsg(value any) string {
	return fmt.Sprintf("expected type %v but got %T", r.expected, value)
}
