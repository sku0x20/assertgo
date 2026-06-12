package assertion

import (
	"github.com/sku0x20/assertgo/pkg/comparator"
	"github.com/sku0x20/assertgo/pkg/matcher"
	"github.com/sku0x20/assertgo/pkg/matcherasserter"
)

type Assertion[V any] struct {
	asserter   *matcherasserter.MatcherAsserter[V]
	comparator comparator.Comparator[V]
}

func New[V any](ma *matcherasserter.MatcherAsserter[V]) *Assertion[V] {
	return &Assertion[V]{
		asserter:   ma,
		comparator: comparator.NewDeepReflectComparator[V](),
	}
}

func (a *Assertion[V]) Not() *Assertion[V] {
	a.asserter.Chain(matcher.NewNotMatcher[V](nil))
	return a
}

func (a *Assertion[V]) EqualTo(other V) *Assertion[V] {
	a.asserter.Assert(matcher.NewEqualMatcher(other, a.comparator))
	return a
}

func (a *Assertion[V]) SameAs(other V) *Assertion[V] {
	a.asserter.Assert(matcher.NewReferenceMatcher(other))
	return a
}

func (a *Assertion[V]) IsNil() *Assertion[V] {
	a.asserter.Assert(matcher.NewNilMatcher[V]())
	return a
}

func (a *Assertion[V]) Matches(m matcher.Matcher[V]) *Assertion[V] {
	a.asserter.Assert(m)
	return a
}
