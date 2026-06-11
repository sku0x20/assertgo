package reference

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/matcher/reference"
)

func Test_SameReferenceMatcher_Match(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		obj := new(int)
		m := reference.NewSameReferenceMatcher(obj)
		if !m.Match(obj) {
			t.Fatal("expected match to be true")
		}
	})
	t.Run("fail", func(t *testing.T) {
		m := reference.NewSameReferenceMatcher(new(int))
		if m.Match(new(int)) {
			t.Fatal("expected match to be false")
		}
	})
}

func Test_SameReferenceMatcher_FailureMsg(t *testing.T) {
	other := new(int)
	value := new(int)
	m := reference.NewSameReferenceMatcher(other)
	msg := m.FailureMsg(value)
	if msg == "" {
		t.Fatal("expected non-empty failure message")
	}
}
