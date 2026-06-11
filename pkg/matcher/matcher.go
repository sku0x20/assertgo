package matcher

type Matcher interface {
	Match() bool
	FailureMsg() string
}
