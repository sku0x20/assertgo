package assertions

import (
	"fmt"
	"reflect"
)

type Sink interface {
	Fail(msg string)
}

type AnyAssertions struct {
	sink  Sink
	value any
}

func NewAnyAssertions(sink Sink, value any) *AnyAssertions {
	return &AnyAssertions{
		sink:  sink,
		value: value,
	}
}

func (a *AnyAssertions) IsEqualTo(other any) {
	if !reflect.DeepEqual(a.value, other) {
		a.sink.Fail(fmt.Sprintf("expected %v but got %v", other, a.value))
	}
}
