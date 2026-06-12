package assertion

import (
	"github.com/sku0x20/assertgo/pkg/comparator"
	"github.com/sku0x20/assertgo/pkg/matcher"
	"github.com/sku0x20/assertgo/pkg/matcherasserter"
	slicematcher "github.com/sku0x20/assertgo/pkg/matcher/slice"
)

type SliceAssertion[E any] struct {
	asserter   *matcherasserter.MatcherAsserter[[]E]
	comparator comparator.Comparator[E]
}

func NewSlice[E any](ma *matcherasserter.MatcherAsserter[[]E]) *SliceAssertion[E] {
	return &SliceAssertion[E]{
		asserter:   ma,
		comparator: comparator.NewDeepReflectComparator[E](),
	}
}

func (a *SliceAssertion[E]) WithComparator(c comparator.Comparator[E]) *SliceAssertion[E] {
	a.comparator = c
	return a
}

func (a *SliceAssertion[E]) HasLength(length int) *SliceAssertion[E] {
	a.asserter.Assert(slicematcher.NewLengthMatcher[E](length))
	return a
}

func (a *SliceAssertion[E]) IsEmpty() *SliceAssertion[E] {
	a.asserter.Assert(slicematcher.NewEmptyMatcher[E]())
	return a
}

func (a *SliceAssertion[E]) EqualTo(other []E) *SliceAssertion[E] {
	a.asserter.Assert(slicematcher.NewEqualMatcher(other, a.comparator))
	return a
}

func (a *SliceAssertion[E]) ContainsAll(elements []E) *SliceAssertion[E] {
	a.asserter.Assert(slicematcher.NewContainsAllMatcher(elements, a.comparator))
	return a
}

func (a *SliceAssertion[E]) ContainsNone(elements []E) *SliceAssertion[E] {
	a.asserter.Assert(matcher.NewNotMatcher[[]E](slicematcher.NewContainsAllMatcher(elements, a.comparator)))
	return a
}
