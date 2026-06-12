package string_matcher

import (
	"fmt"
	"regexp"
)

type RegexMatcher struct {
	pattern *regexp.Regexp
}

func NewRegexMatcher(pattern string) *RegexMatcher {
	return &RegexMatcher{pattern: regexp.MustCompile(pattern)}
}

func (r *RegexMatcher) Match(value string) bool {
	return r.pattern.MatchString(value)
}

func (r *RegexMatcher) FailureMsg(value string) string {
	return fmt.Sprintf("expected %q to match pattern %q", value, r.pattern)
}
