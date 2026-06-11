package matcher

type Matcher[T any] interface {
	Match(value T) bool
	FailureMsg(value T) string
}

type LazyMatcher[T any, E any] interface {
	Matcher[T]
	Set(v E)
}
