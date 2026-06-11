package test

type MockMatcher struct {
	MatchResult bool
	Msg         string
}

func (m *MockMatcher) Match(_ any) bool {
	return m.MatchResult
}

func (m *MockMatcher) FailureMsg(_ any) string {
	return m.Msg
}
