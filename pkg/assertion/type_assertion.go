package assertion

import (
	"reflect"

	"github.com/sku0x20/assertgo/pkg/matcher"
	"github.com/sku0x20/assertgo/pkg/matcherasserter"
)

type TypeAssertion struct {
	asserter *matcherasserter.MatcherAsserter[any]
}

func NewType(ma *matcherasserter.MatcherAsserter[any]) *TypeAssertion {
	return &TypeAssertion{asserter: ma}
}

func (a *TypeAssertion) IsType(expected reflect.Type) *TypeAssertion {
	a.asserter.Assert(matcher.NewReflectTypeMatcher(expected))
	return a
}

// todo: replace with Is[T any]() method once Go supports generic methods
func Is[T any](a *TypeAssertion) *TypeAssertion {
	return a.IsType(reflect.TypeOf(*new(T)))
}
