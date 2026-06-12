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

// todo: use generic method in golang 1.27

func (a *AssertionSuit) Assert(value any) *assertion.Assertion[any] {
	return assertion.New(matcherasserter.New(a.sink, value))
}

func (a *AssertionSuit) AssertSlice(value []any) *assertion.SliceAssertion[any] {
	return assertion.NewSlice[any](matcherasserter.New(a.sink, value))
}
