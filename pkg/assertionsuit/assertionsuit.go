package assertionsuit

import (
	"github.com/sku0x20/assertgo/pkg/assertion"
	"github.com/sku0x20/assertgo/pkg/constraints"
	"github.com/sku0x20/assertgo/pkg/matcherasserter"
	"github.com/sku0x20/assertgo/pkg/sink"
)

type AssertionSuit struct {
	sink *sink.TSink
}

func New(s *sink.TSink) *AssertionSuit {
	return &AssertionSuit{sink: s}
}

func (a *AssertionSuit) Assert[T any](value T) *assertion.Assertion[T] {
	return assertion.New(matcherasserter.New(a.sink, value))
}

func (a *AssertionSuit) AssertSlice(value []any) *assertion.SliceAssertion[any] {
	return assertion.NewSlice[any](matcherasserter.New(a.sink, value))
}

func (a *AssertionSuit) AssertBool(value bool) *assertion.BoolAssertion {
	return assertion.NewBool(matcherasserter.New(a.sink, value))
}

func (a *AssertionSuit) AssertInt[T constraints.Integer](value T) *assertion.NumberAssertion[T] {
	return assertion.NewNumber[T](matcherasserter.New(a.sink, value))
}

// todo: use generic method in golang 1.27
func (a *AssertionSuit) AssertFloat(value float64) *assertion.NumberAssertion[float64] {
	return assertion.NewNumber[float64](matcherasserter.New(a.sink, value))
}

// todo: use generic method in golang 1.27
func (a *AssertionSuit) AssertMap(value map[any]any) *assertion.MapAssertion[any, any] {
	return assertion.NewMap[any, any](matcherasserter.New(a.sink, value))
}

func (a *AssertionSuit) AssertString(value string) *assertion.StringAssertion {
	return assertion.NewString(matcherasserter.New(a.sink, value))
}

func (a *AssertionSuit) AssertType(value any) *assertion.TypeAssertion {
	return assertion.NewType(matcherasserter.New(a.sink, value))
}
