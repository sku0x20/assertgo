package comparator

import "reflect"

type ReferenceComparator[V any] struct{}

func NewReferenceComparator[V any]() *ReferenceComparator[V] {
	return &ReferenceComparator[V]{}
}

func (r *ReferenceComparator[V]) Compare(a, b V) int {
	if reflect.ValueOf(a).Pointer() == reflect.ValueOf(b).Pointer() {
		return 0
	}
	return -1
}
