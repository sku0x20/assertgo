package matcher

type LazyExpected[T any] interface {
	Set(v Matcher[T])
}
