package assertion

import (
	"github.com/sku0x20/assertgo/pkg/comparator"
	"github.com/sku0x20/assertgo/pkg/matcher"
	"github.com/sku0x20/assertgo/pkg/matcherasserter"
	"github.com/sku0x20/assertgo/pkg/matcher/collection"
	string_matcher "github.com/sku0x20/assertgo/pkg/matcher/string"
)

type StringAssertion struct {
	asserter   *matcherasserter.MatcherAsserter[string]
	comparator comparator.Comparator[string]
}

func NewString(ma *matcherasserter.MatcherAsserter[string]) *StringAssertion {
	return &StringAssertion{
		asserter:   ma,
		comparator: comparator.NewStringComparator(),
	}
}

func (a *StringAssertion) Not() *StringAssertion {
	a.asserter.Chain(matcher.NewNotMatcher[string](nil))
	return a
}

func (a *StringAssertion) EqualTo(other string) *StringAssertion {
	a.asserter.Assert(matcher.NewEqualMatcher(other, a.comparator))
	return a
}

func (a *StringAssertion) Contains(substr string) *StringAssertion {
	a.asserter.Assert(string_matcher.NewContainsMatcher(substr))
	return a
}

func (a *StringAssertion) HasPrefix(prefix string) *StringAssertion {
	a.asserter.Assert(string_matcher.NewHasPrefixMatcher(prefix))
	return a
}

func (a *StringAssertion) HasSuffix(suffix string) *StringAssertion {
	a.asserter.Assert(string_matcher.NewHasSuffixMatcher(suffix))
	return a
}

func (a *StringAssertion) HasLength(length int) *StringAssertion {
	a.asserter.Assert(collection.NewLengthMatcher[string](length))
	return a
}

func (a *StringAssertion) IsEmpty() *StringAssertion {
	return a.HasLength(0)
}

func (a *StringAssertion) MatchesRegex(pattern string) *StringAssertion {
	a.asserter.Assert(string_matcher.NewRegexMatcher(pattern))
	return a
}
