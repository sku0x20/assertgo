package test

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg"
)

func Test_MatcherAsserter_pass(t *testing.T) {
	mock, sink := NewSink()
	pkg.MatcherAsserter(sink, &MockMatcher{MatchResult: true}, nil)
	if mock.FatalCalled {
		t.Fatal("expected no failure")
	}
}

func Test_MatcherAsserter_fail(t *testing.T) {
	mock, sink := NewSink()
	pkg.MatcherAsserter(sink, &MockMatcher{MatchResult: false, Msg: "failed"}, nil)
	if !mock.FatalCalled {
		t.Fatal("expected failure")
	}
}
