package comparator

type Comparator[V any] interface {
	Equal(a, b V) bool
}
