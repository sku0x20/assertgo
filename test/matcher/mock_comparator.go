package matcher

type MockComparator[V any] struct {
	Result int
}

func (m *MockComparator[V]) Compare(_, _ V) int {
	return m.Result
}
