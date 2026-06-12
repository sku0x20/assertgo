package string_matcher

import (
	"fmt"
	"strings"
)

type ContainsMatcher struct {
	substr string
}

func NewContainsMatcher(substr string) *ContainsMatcher {
	return &ContainsMatcher{substr: substr}
}

func (c *ContainsMatcher) Match(value string) bool {
	return strings.Contains(value, c.substr)
}

func (c *ContainsMatcher) FailureMsg(value string) string {
	return fmt.Sprintf("expected %q to contain %q", value, c.substr)
}
