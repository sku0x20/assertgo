package matcherasserter

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/matcherasserter"
	testmatcher "github.com/sku0x20/assertgo/test/matcher"
	agtest "github.com/sku0x20/assertgo/test"
)

func Test_MatcherAsserter_pass(t *testing.T) {
	mock, sink := agtest.NewSink()
	matcherasserter.New(sink, nil).Assert(&testmatcher.MockMatcher{MatchResult: true})
	if mock.FatalCalled {
		t.Fatal("expected no failure")
	}
}

func Test_MatcherAsserter_fail(t *testing.T) {
	mock, sink := agtest.NewSink()
	matcherasserter.New(sink, nil).Assert(&testmatcher.MockMatcher{MatchResult: false, Msg: "failed"})
	if !mock.FatalCalled {
		t.Fatal("expected failure")
	}
}
