package reference

import (
	"fmt"

	"github.com/sku0x20/assertgo/pkg/matcher"
)

var _ matcher.Matcher[any] = (*SameReferenceMatcher)(nil)

type SameReferenceMatcher struct {
	other any
}

func NewSameReferenceMatcher(other any) *SameReferenceMatcher {
	return &SameReferenceMatcher{other: other}
}

func (s *SameReferenceMatcher) Match(value any) bool {
	return value == s.other
}

func (s *SameReferenceMatcher) FailureMsg(value any) string {
	return fmt.Sprintf("expected same reference as %v but got %v", s.other, value)
}
