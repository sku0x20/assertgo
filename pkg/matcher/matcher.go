package matcher

type Matcher interface {
	Match(value any) bool
	FailureMsg(value any) string
}
