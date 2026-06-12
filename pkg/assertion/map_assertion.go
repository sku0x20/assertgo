package assertion

import (
	"github.com/sku0x20/assertgo/pkg/comparator"
	"github.com/sku0x20/assertgo/pkg/matcher"
	"github.com/sku0x20/assertgo/pkg/matcherasserter"
	map_matcher "github.com/sku0x20/assertgo/pkg/matcher/map"
)

type MapAssertion[K comparable, V any] struct {
	asserter   *matcherasserter.MatcherAsserter[map[K]V]
	comparator comparator.Comparator[V]
}

func NewMap[K comparable, V any](ma *matcherasserter.MatcherAsserter[map[K]V]) *MapAssertion[K, V] {
	return &MapAssertion[K, V]{
		asserter:   ma,
		comparator: comparator.NewDeepReflectComparator[V](),
	}
}

func (a *MapAssertion[K, V]) Not() *MapAssertion[K, V] {
	a.asserter.Chain(matcher.NewNotMatcher[map[K]V](nil))
	return a
}

func (a *MapAssertion[K, V]) WithComparator(c comparator.Comparator[V]) *MapAssertion[K, V] {
	a.comparator = c
	return a
}

func (a *MapAssertion[K, V]) EqualTo(other map[K]V) *MapAssertion[K, V] {
	a.asserter.Assert(map_matcher.NewEqualMatcher(other, a.comparator))
	return a
}

func (a *MapAssertion[K, V]) ContainsAllEntries(entries map[K]V) *MapAssertion[K, V] {
	a.asserter.Assert(map_matcher.NewContainsAllEntriesMatcher(entries, a.comparator))
	return a
}

func (a *MapAssertion[K, V]) HasLength(length int) *MapAssertion[K, V] {
	a.asserter.Assert(map_matcher.NewLengthMatcher[K, V](length))
	return a
}

func (a *MapAssertion[K, V]) IsEmpty() *MapAssertion[K, V] {
	return a.HasLength(0)
}
