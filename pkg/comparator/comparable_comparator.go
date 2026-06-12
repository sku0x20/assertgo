package comparator

type ComparableComparator[V comparable] struct{}

func NewComparableComparator[V comparable]() *ComparableComparator[V] {
	return &ComparableComparator[V]{}
}

func (c *ComparableComparator[V]) Compare(a, b V) int {
	if a == b {
		return 0
	}
	return -1
}
