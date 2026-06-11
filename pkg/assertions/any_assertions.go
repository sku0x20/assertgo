package assertions

import "github.com/sku0x20/assertgo/pkg"

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

func IsEqualTo(other any) {
	// todo: fill in
}
