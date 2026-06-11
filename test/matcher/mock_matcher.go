package matcher

type mockMatcher struct {
	matchResult bool
	failureMsg  string
}

func (m *mockMatcher) Match(_ any) bool {
	return m.matchResult
}

func (m *mockMatcher) FailureMsg(_ any) string {
	return m.failureMsg
}
