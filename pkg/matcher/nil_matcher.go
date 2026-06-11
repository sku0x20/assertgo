package matcher

import "reflect"

type NilMatcher[V any] struct{}

func NewNilMatcher[V any]() *NilMatcher[V] {
	return &NilMatcher[V]{}
}

func (n *NilMatcher[V]) Match(value V) bool {
	if any(value) == nil {
		return true
	}
	rv := reflect.ValueOf(any(value))
	switch rv.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func:
		return rv.IsNil()
	}
	return false
}

func (n *NilMatcher[V]) FailureMsg(value V) string {
	return "expected nil"
}
