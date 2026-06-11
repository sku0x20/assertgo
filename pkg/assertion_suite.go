package pkg

import "github.com/sku0x20/assertgo/pkg/assertions"

type AssertionSuite struct {
	sink *TSink
}

func NewAssertionSuite(sink *TSink) *AssertionSuite {
	return &AssertionSuite{sink: sink}
}

func (a *AssertionSuite) Assert(value any) *assertions.AnyAssertions {
	return assertions.NewAnyAssertions(a.sink, value)
}
