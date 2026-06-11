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
