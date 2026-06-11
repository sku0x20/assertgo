package matcher

import "fmt"

type ReferenceMatcher struct {
	other any
}

func NewReferenceMatcher(other any) *ReferenceMatcher {
	return &ReferenceMatcher{other: other}
}

func (r *ReferenceMatcher) Match(value any) bool {
	return value == r.other
}

func (r *ReferenceMatcher) FailureMsg(value any) string {
	return fmt.Sprintf("expected same reference as %v but got %v", r.other, value)
}
