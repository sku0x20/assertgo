package matcher

type NotMatcher[T any] struct {
	matcher Matcher[T]
}

func NewNotMatcher[T any](m Matcher[T]) *NotMatcher[T] {
	return &NotMatcher[T]{matcher: m}
}

func (n *NotMatcher[T]) Match(value T) bool {
	return !n.matcher.Match(value)
}

func (n *NotMatcher[T]) FailureMsg(value T) string {
	return "not: " + n.matcher.FailureMsg(value)
}
