package test

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg"
	testmatcher "github.com/sku0x20/assertgo/test/matcher"
)

func Test_MatcherAsserter_pass(t *testing.T) {
	mock, sink := NewSink()
	pkg.NewMatcherAsserter(sink, nil).Assert(&testmatcher.MockMatcher{MatchResult: true})
	if mock.FatalCalled {
		t.Fatal("expected no failure")
	}
}

func Test_MatcherAsserter_fail(t *testing.T) {
	mock, sink := NewSink()
	pkg.NewMatcherAsserter(sink, nil).Assert(&testmatcher.MockMatcher{MatchResult: false, Msg: "failed"})
	if !mock.FatalCalled {
		t.Fatal("expected failure")
	}
}
