package reference

import "fmt"

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
