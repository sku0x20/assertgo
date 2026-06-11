package assertions

import (
	"fmt"
	"reflect"

	"github.com/sku0x20/assertgo/pkg/sink"
)

type AnyAssertions struct {
	sink  *sink.TSink
	value any
}

func NewAnyAssertions(s *sink.TSink, value any) *AnyAssertions {
	return &AnyAssertions{
		sink:  s,
		value: value,
	}
}

func (a *AnyAssertions) IsEqualTo(other any) {
	if !reflect.DeepEqual(a.value, other) {
		a.sink.Fail(fmt.Sprintf("expected %v but got %v", other, a.value))
	}
}
