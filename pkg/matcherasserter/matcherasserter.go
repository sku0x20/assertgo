package matcherasserter

import (
	"github.com/sku0x20/assertgo/pkg/matcher"
	"github.com/sku0x20/assertgo/pkg/sink"
)

type MatcherAsserter struct {
	sink            *sink.TSink
	value           any
	chain matcher.Matcher[any]
}

func New(s *sink.TSink, value any) *MatcherAsserter {
	return &MatcherAsserter{sink: s, value: value}
}

func (ma *MatcherAsserter) Chain(m matcher.Matcher[any]) {
	ma.chain = m
}

func (ma *MatcherAsserter) Assert(m matcher.Matcher[any]) {
	if !m.Match(ma.value) {
		ma.sink.Fail(m.FailureMsg(ma.value))
	}
}
