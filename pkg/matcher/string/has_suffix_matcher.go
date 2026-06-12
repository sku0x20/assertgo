package string_matcher

import (
	"fmt"
	"strings"
)

type HasSuffixMatcher struct {
	suffix string
}

func NewHasSuffixMatcher(suffix string) *HasSuffixMatcher {
	return &HasSuffixMatcher{suffix: suffix}
}

func (h *HasSuffixMatcher) Match(value string) bool {
	return strings.HasSuffix(value, h.suffix)
}

func (h *HasSuffixMatcher) FailureMsg(value string) string {
	return fmt.Sprintf("expected %q to have suffix %q", value, h.suffix)
}
