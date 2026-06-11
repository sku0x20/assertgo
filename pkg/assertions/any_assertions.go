package assertions

import (
	"fmt"
	"reflect"

	"github.com/sku0x20/assertgo/pkg"
)

type AnyAssertions struct {
	sink  *pkg.TSink
	value any
}

func NewAnyAssertions(sink *pkg.TSink, value any) *AnyAssertions {
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
