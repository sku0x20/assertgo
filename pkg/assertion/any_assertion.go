package assertion

import (
	"github.com/sku0x20/assertgo/pkg/matcher"
	"github.com/sku0x20/assertgo/pkg/matcherasserter"
)

type AnyAssertion struct {
	asserter *matcherasserter.MatcherAsserter[any]
}

func NewAnyAssertion(ma *matcherasserter.MatcherAsserter[any]) *AnyAssertion {
	return &AnyAssertion{asserter: ma}
}

func (a *AnyAssertion) EqualTo(other any) *AnyAssertion {
	a.asserter.Assert(matcher.NewDeepEqualMatcher(other))
	return a
}

func (a *AnyAssertion) SameAs(other any) *AnyAssertion {
	a.asserter.Assert(matcher.NewReferenceMatcher(other))
	return a
}

func (a *AnyAssertion) NotEqualTo(other any) *AnyAssertion {
	a.asserter.Chain(matcher.NewNotMatcher[any](nil))
	a.asserter.Assert(matcher.NewDeepEqualMatcher(other))
	return a
}

func (a *AnyAssertion) NotSameAs(other any) *AnyAssertion {
	a.asserter.Chain(matcher.NewNotMatcher[any](nil))
	a.asserter.Assert(matcher.NewReferenceMatcher(other))
	return a
}

func (a *AnyAssertion) IsNil() *AnyAssertion {
	a.asserter.Assert(matcher.NewNilMatcher())
	return a
}

func (a *AnyAssertion) IsNotNil() *AnyAssertion {
	a.asserter.Chain(matcher.NewNotMatcher[any](nil))
	a.asserter.Assert(matcher.NewNilMatcher())
	return a
}
