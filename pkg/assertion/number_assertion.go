package assertion

import (
	"github.com/sku0x20/assertgo/pkg/comparator"
	"github.com/sku0x20/assertgo/pkg/constraints"
	"github.com/sku0x20/assertgo/pkg/matcher"
	"github.com/sku0x20/assertgo/pkg/matcherasserter"
	numbermatcher "github.com/sku0x20/assertgo/pkg/matcher/number"
)

type NumberAssertion[V constraints.Number] struct {
	asserter   *matcherasserter.MatcherAsserter[V]
	comparator comparator.Comparator[V]
}

func NewNumber[V constraints.Number](ma *matcherasserter.MatcherAsserter[V]) *NumberAssertion[V] {
	return &NumberAssertion[V]{
		asserter:   ma,
		comparator: comparator.NewNumberComparator[V](),
	}
}

func (a *NumberAssertion[V]) Not() *NumberAssertion[V] {
	a.asserter.Chain(matcher.NewNotMatcher[V](nil))
	return a
}

func (a *NumberAssertion[V]) WithComparator(c comparator.Comparator[V]) *NumberAssertion[V] {
	a.comparator = c
	return a
}

func (a *NumberAssertion[V]) EqualTo(other V) *NumberAssertion[V] {
	a.asserter.Assert(matcher.NewEqualMatcher(other, a.comparator))
	return a
}

func (a *NumberAssertion[V]) GreaterThan(other V) *NumberAssertion[V] {
	a.asserter.Assert(numbermatcher.NewGreaterThanMatcher(other, a.comparator))
	return a
}

func (a *NumberAssertion[V]) GreaterThanOrEqual(other V) *NumberAssertion[V] {
	a.asserter.Assert(numbermatcher.NewGreaterThanOrEqualMatcher(other, a.comparator))
	return a
}

func (a *NumberAssertion[V]) LessThan(other V) *NumberAssertion[V] {
	a.asserter.Assert(numbermatcher.NewLessThanMatcher(other, a.comparator))
	return a
}

func (a *NumberAssertion[V]) LessThanOrEqual(other V) *NumberAssertion[V] {
	a.asserter.Assert(numbermatcher.NewLessThanOrEqualMatcher(other, a.comparator))
	return a
}

func (a *NumberAssertion[V]) InRange(min, max V) *NumberAssertion[V] {
	a.asserter.Assert(numbermatcher.NewInRangeMatcher(min, max, a.comparator))
	return a
}
