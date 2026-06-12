package string_matcher

import (
	"fmt"
	"strings"
)

type HasPrefixMatcher struct {
	prefix string
}

func NewHasPrefixMatcher(prefix string) *HasPrefixMatcher {
	return &HasPrefixMatcher{prefix: prefix}
}

func (h *HasPrefixMatcher) Match(value string) bool {
	return strings.HasPrefix(value, h.prefix)
}

func (h *HasPrefixMatcher) FailureMsg(value string) string {
	return fmt.Sprintf("expected %q to have prefix %q", value, h.prefix)
}
