package matcher

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/matcher"
)

func Test_ReferenceMatcher_Match(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		obj := new(int)
		m := matcher.NewReferenceMatcher(obj)
		if !m.Match(obj) {
			t.Fatal("expected match to be true")
		}
	})
	t.Run("fail", func(t *testing.T) {
		m := matcher.NewReferenceMatcher(new(int))
		if m.Match(new(int)) {
			t.Fatal("expected match to be false")
		}
	})
}

func Test_ReferenceMatcher_FailureMsg(t *testing.T) {
	obj := new(int)
	m := matcher.NewReferenceMatcher(obj)
	msg := m.FailureMsg(new(int))
	if msg == "" {
		t.Fatal("expected non-empty failure message")
	}
}
