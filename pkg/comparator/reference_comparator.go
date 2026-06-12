package comparator

type ReferenceComparator[V any] struct{}

func NewReferenceComparator[V any]() *ReferenceComparator[V] {
	return &ReferenceComparator[V]{}
}

func (r *ReferenceComparator[V]) Compare(a, b V) int {
	if any(a) == any(b) {
		return 0
	}
	return -1
}
