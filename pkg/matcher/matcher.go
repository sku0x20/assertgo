package matcher

type Matcher[T any] interface {
	Match(value T) bool
	FailureMsg(value T) string
}
