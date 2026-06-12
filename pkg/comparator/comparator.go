package comparator

type Comparator[V any] interface {
	// Compare returns -1 if a < b, 0 if a == b, 1 if a > b.
	Compare(a, b V) int
}
