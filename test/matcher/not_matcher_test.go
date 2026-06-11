package matcher

import (
	"testing"

	"github.com/sku0x20/assertgo/pkg/matcher"
	"github.com/sku0x20/assertgo/pkg/matcher/reference"
)

func Test_NotMatcher_Match(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		m := matcher.NewNotMatcher(reference.NewSameReferenceMatcher(new(int)))
		if !m.Match(new(int)) {
			t.Fatal("expected match to be true")
		}
	})
	t.Run("fail", func(t *testing.T) {
		obj := new(int)
		m := matcher.NewNotMatcher(reference.NewSameReferenceMatcher(obj))
		if m.Match(obj) {
			t.Fatal("expected match to be false")
		}
	})
}
