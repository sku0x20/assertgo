package assertionsuit

import (
	"github.com/sku0x20/assertgo/pkg/assertion"
	"github.com/sku0x20/assertgo/pkg/matcherasserter"
	"github.com/sku0x20/assertgo/pkg/sink"
)

type AssertionSuit struct {
	sink *sink.TSink
}

func New(s *sink.TSink) *AssertionSuit {
	return &AssertionSuit{sink: s}
}

func (a *AssertionSuit) Assert(value any) *assertion.AnyAssertion {
	return assertion.NewAnyAssertion(matcherasserter.New(a.sink, value))
}
