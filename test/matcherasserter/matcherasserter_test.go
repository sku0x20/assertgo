package matcherasserter

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/matcher"
	"github.com/sku0x20/assertgo/pkg/matcherasserter"
	agtest "github.com/sku0x20/assertgo/test"
	testmatcher "github.com/sku0x20/assertgo/test/matcher"
)

func Test_MatcherAsserter_pass(t *testing.T) {
	mock, sink := agtest.NewSink()
	matcherasserter.New[any](sink, nil).Assert(&testmatcher.MockMatcher{MatchResult: true})
	if mock.FatalCalled {
		t.Fatal("expected no failure")
	}
}

func Test_MatcherAsserter_fail(t *testing.T) {
	mock, sink := agtest.NewSink()
	matcherasserter.New[any](sink, nil).Assert(&testmatcher.MockMatcher{MatchResult: false, Msg: "failed"})
	if !mock.FatalCalled {
		t.Fatal("expected failure")
	}
}

func Test_MatcherAsserter_Chain(t *testing.T) {
	mock, sink := agtest.NewSink()
	ma := matcherasserter.New[any](sink, nil)
	ma.Chain(matcher.NewNotMatcher[any](nil))
	ma.Assert(&testmatcher.MockMatcher{MatchResult: true})
	if !mock.FatalCalled {
		t.Fatal("expected failure: NotMatcher should negate the passing matcher")
	}
	mock.FatalCalled = false
	ma.Assert(&testmatcher.MockMatcher{MatchResult: true})
	if mock.FatalCalled {
		t.Fatal("expected no failure: chain should be cleared after first assert")
	}
}
