package data

type Queue[T any] struct {
	q []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{make([]T, 0)}
}

func (q *Queue[T]) Enqueue(v T) {
	q.q = append(q.q, v)
}

func (q *Queue[T]) Peek() T {
	return q.q[0]
}

func (q *Queue[T]) Dequeue() T {
	v := q.Peek()
	q.q = q.q[1:]
	return v
}

func (q *Queue[T]) Len() int {
	return len(q.q)
}

func (q *Queue[T]) Empty() bool {
	return q.Len() == 0
}
