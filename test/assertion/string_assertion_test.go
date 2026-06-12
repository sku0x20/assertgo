package assertion

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/assertion"
	"github.com/sku0x20/assertgo/pkg/matcherasserter"
	agtest "github.com/sku0x20/assertgo/test"
)

func Test_StringAssertion_EqualTo(t *testing.T) {
	assertStringPasses(t, "hello", func(a *assertion.StringAssertion) { a.EqualTo("hello") })
}

func Test_StringAssertion_Contains(t *testing.T) {
	assertStringPasses(t, "hello", func(a *assertion.StringAssertion) { a.Contains("ell") })
}

func Test_StringAssertion_HasPrefix(t *testing.T) {
	assertStringPasses(t, "hello", func(a *assertion.StringAssertion) { a.HasPrefix("hel") })
}

func Test_StringAssertion_HasSuffix(t *testing.T) {
	assertStringPasses(t, "hello", func(a *assertion.StringAssertion) { a.HasSuffix("llo") })
}

func Test_StringAssertion_HasLength(t *testing.T) {
	assertStringPasses(t, "hello", func(a *assertion.StringAssertion) { a.HasLength(5) })
}

func Test_StringAssertion_IsEmpty(t *testing.T) {
	assertStringPasses(t, "", func(a *assertion.StringAssertion) { a.IsEmpty() })
}

func Test_StringAssertion_MatchesRegex(t *testing.T) {
	assertStringPasses(t, "hello123", func(a *assertion.StringAssertion) { a.MatchesRegex(`^[a-z]+\d+$`) })
}

func Test_StringAssertion_Not(t *testing.T) {
	assertStringPasses(t, "hello", func(a *assertion.StringAssertion) { a.Not().EqualTo("world") })
}

func assertStringPasses(t *testing.T, value string, fn func(*assertion.StringAssertion)) {
	t.Helper()
	mock, s := agtest.NewSink()
	a := assertion.NewString(matcherasserter.New(s, value))
	fn(a)
	if mock.FatalCalled {
		t.Fatal("expected no failure")
	}
}
