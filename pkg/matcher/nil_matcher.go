package matcher

import "reflect"

type NilMatcher struct{}

func NewNilMatcher() *NilMatcher {
	return &NilMatcher{}
}

func (n *NilMatcher) Match(value any) bool {
	if value == nil {
		return true
	}
	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func:
		return rv.IsNil()
	}
	return false
}

func (n *NilMatcher) FailureMsg(value any) string {
	return "expected nil but got non-nil"
}
