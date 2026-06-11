package matcher

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg"
	agtest "github.com/sku0x20/assertgo/test"
)

func Test_MatcherAsserter_pass(t *testing.T) {
	mock, sink := agtest.NewSink()
	pkg.MatcherAsserter(sink, &mockMatcher{matchResult: true}, nil)
	if mock.FatalCalled {
		t.Fatal("expected no failure")
	}
}

func Test_MatcherAsserter_fail(t *testing.T) {
	mock, sink := agtest.NewSink()
	pkg.MatcherAsserter(sink, &mockMatcher{matchResult: false, failureMsg: "failed"}, nil)
	if !mock.FatalCalled {
		t.Fatal("expected failure")
	}
}
