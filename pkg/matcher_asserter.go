package pkg

import (
	"github.com/sku0x20/assertgo/pkg/matcher"
	"github.com/sku0x20/assertgo/pkg/sink"
)

func MatcherAsserter[T any](s *sink.TSink, m matcher.Matcher[T], value T) {
	if !m.Match(value) {
		s.Fail(m.FailureMsg(value))
	}
}
