package equal

type Comparator[V any] interface {
	Equal(a, b V) bool
}
