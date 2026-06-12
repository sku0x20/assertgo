package comparator

import "reflect"

type DeepReflectComparator[V any] struct{}

func NewDeepReflectComparator[V any]() *DeepReflectComparator[V] {
	return &DeepReflectComparator[V]{}
}

func (d *DeepReflectComparator[V]) Compare(a, b V) int {
	if reflect.DeepEqual(a, b) {
		return 0
	}
	return -1
}
