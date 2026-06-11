package assertion

import (
	"fmt"
	"reflect"

	"github.com/sku0x20/assertgo/pkg/sink"
)

type AnyAssertion struct {
	sink  *sink.TSink
	value any
}

func NewAnyAssertion(s *sink.TSink, value any) *AnyAssertion {
	return &AnyAssertion{
		sink:  s,
		value: value,
	}
}

func (a *AnyAssertion) IsEqualTo(other any) *AnyAssertion {
	if !reflect.DeepEqual(a.value, other) {
		a.sink.Fail(fmt.Sprintf("expected %v but got %v", other, a.value))
	}
	return a
}

func (a *AnyAssertion) IsSameAs(other any) *AnyAssertion {
	if a.value != other {
		a.sink.Fail(fmt.Sprintf("expected same reference as %v but got %v", other, a.value))
	}
	return a
}

func (a *AnyAssertion) IsNotEqualTo(other any) *AnyAssertion {
	if reflect.DeepEqual(a.value, other) {
		a.sink.Fail(fmt.Sprintf("expected %v to not equal %v", a.value, other))
	}
	return a
}

func (a *AnyAssertion) IsNotSameAs(other any) *AnyAssertion {
	if a.value == other {
		a.sink.Fail(fmt.Sprintf("expected different reference but got same %v", a.value))
	}
	return a
}

func (a *AnyAssertion) IsNil() *AnyAssertion {
	if !isNil(a.value) {
		a.sink.Fail(fmt.Sprintf("expected nil but got %v", a.value))
	}
	return a
}

func (a *AnyAssertion) IsNotNil() *AnyAssertion {
	if isNil(a.value) {
		a.sink.Fail(fmt.Sprintf("expected not nil but got nil"))
	}
	return a
}

func isNil(v any) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func:
		return rv.IsNil()
	}
	return false
}
