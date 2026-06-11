package pkg

import (
	"github.com/sku0x20/assertgo/pkg/assertion"
	"github.com/sku0x20/assertgo/pkg/sink"
)

type AssertionSuite struct {
	sink *sink.TSink
}

func NewAssertionSuite(s *sink.TSink) *AssertionSuite {
	return &AssertionSuite{sink: s}
}

func (a *AssertionSuite) Assert(value any) *assertion.AnyAssertion {
	return assertion.NewAnyAssertion(a.sink, value)
}
