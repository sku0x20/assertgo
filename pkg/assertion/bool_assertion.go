package assertion

import (
	"github.com/sku0x20/assertgo/pkg/matcher"
	"github.com/sku0x20/assertgo/pkg/matcherasserter"
)

type BoolAssertion struct {
	asserter *matcherasserter.MatcherAsserter[bool]
}

func NewBool(ma *matcherasserter.MatcherAsserter[bool]) *BoolAssertion {
	return &BoolAssertion{asserter: ma}
}

func (a *BoolAssertion) Not() *BoolAssertion {
	a.asserter.Chain(matcher.NewNotMatcher[bool](nil))
	return a
}

func (a *BoolAssertion) IsTrue() *BoolAssertion {
	a.asserter.Assert(matcher.NewBoolMatcher(true))
	return a
}

func (a *BoolAssertion) IsFalse() *BoolAssertion {
	a.asserter.Assert(matcher.NewBoolMatcher(false))
	return a
}
