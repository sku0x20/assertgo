package matcherasserter

import (
	"github.com/sku0x20/assertgo/pkg/matcher"
	"github.com/sku0x20/assertgo/pkg/sink"
)

type lazyMatcher[T any] interface {
	matcher.Matcher[T]
	matcher.LazyExpected[T]
}

type MatcherAsserter[T any] struct {
	sink  *sink.TSink
	value T
	chain lazyMatcher[T]
}

func New[T any](s *sink.TSink, value T) *MatcherAsserter[T] {
	return &MatcherAsserter[T]{sink: s, value: value}
}

func (ma *MatcherAsserter[T]) Chain(m lazyMatcher[T]) {
	ma.chain = m
}

func (ma *MatcherAsserter[T]) Assert(m matcher.Matcher[T]) {
	if ma.chain != nil {
		ma.chain.Set(m)
		m = ma.chain
		ma.chain = nil
	}
	if !m.Match(ma.value) {
		ma.sink.Fail(m.FailureMsg(ma.value))
	}
}
