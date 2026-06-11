package assertion

import (
	"github.com/sku0x20/assertgo/pkg/matcher"
	"github.com/sku0x20/assertgo/pkg/matcherasserter"
)

type AnyAssertion struct {
	asserter *matcherasserter.MatcherAsserter
}

func NewAnyAssertion(ma *matcherasserter.MatcherAsserter) *AnyAssertion {
	return &AnyAssertion{asserter: ma}
}

func (a *AnyAssertion) IsEqualTo(other any) *AnyAssertion {
	a.asserter.Assert(matcher.NewDeepEqualMatcher(other))
	return a
}

func (a *AnyAssertion) IsSameAs(other any) *AnyAssertion {
	a.asserter.Assert(matcher.NewReferenceMatcher(other))
	return a
}

func (a *AnyAssertion) IsNotEqualTo(other any) *AnyAssertion {
	a.asserter.Assert(matcher.NewNotMatcher[any](matcher.NewDeepEqualMatcher(other)))
	return a
}

func (a *AnyAssertion) IsNotSameAs(other any) *AnyAssertion {
	a.asserter.Assert(matcher.NewNotMatcher[any](matcher.NewReferenceMatcher(other)))
	return a
}

func (a *AnyAssertion) IsNil() *AnyAssertion {
	a.asserter.Assert(matcher.NewNilMatcher())
	return a
}

func (a *AnyAssertion) IsNotNil() *AnyAssertion {
	a.asserter.Assert(matcher.NewNotMatcher[any](matcher.NewNilMatcher()))
	return a
}
