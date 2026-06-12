package matcher

import "fmt"

type BoolMatcher struct {
	expected bool
}

func NewBoolMatcher(expected bool) *BoolMatcher {
	return &BoolMatcher{expected: expected}
}

func (b *BoolMatcher) Match(value bool) bool {
	return value == b.expected
}

func (b *BoolMatcher) FailureMsg(_ bool) string {
	return fmt.Sprintf("expected %v", b.expected)
}
