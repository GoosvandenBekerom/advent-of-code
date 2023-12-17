package data

type LessThanComparer[T any] interface {
	LessThan(other T) bool
}

type Heap[T LessThanComparer[T]] []T

func (h Heap[T]) Len() int           { return len(h) }
func (h Heap[T]) Less(i, j int) bool { return h[i].LessThan(h[j]) }
func (h Heap[T]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push uses a pointer receiver because it modifies
// the slice's length, not just its contents.
func (h *Heap[T]) Push(x any) {
	*h = append(*h, x.(T))
}

// Pop uses a pointer receiver because it modifies
// the slice's length, not just its contents.
func (h *Heap[T]) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
