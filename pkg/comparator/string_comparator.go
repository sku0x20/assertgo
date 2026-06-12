package comparator

type StringComparator struct{}

func NewStringComparator() *StringComparator {
	return &StringComparator{}
}

func (c *StringComparator) Compare(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}
