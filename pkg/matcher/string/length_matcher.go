package string_matcher

import "fmt"

type LengthMatcher struct {
	length int
}

func NewLengthMatcher(length int) *LengthMatcher {
	return &LengthMatcher{length: length}
}

func (l *LengthMatcher) Match(value string) bool {
	return len(value) == l.length
}

func (l *LengthMatcher) FailureMsg(value string) string {
	return fmt.Sprintf("expected length %d but got %d", l.length, len(value))
}
