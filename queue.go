package jqueue

type Queue[T any] struct {
	head, tail *Chunk[T]
	chunkSize  int
}

func New[T any](chunkSize int) *Queue[T] {
	return &Queue[T]{
		chunkSize: chunkSize,
	}
}

func (q *Queue[T]) initialize() {
	if q.head != nil {
		return
	}

	q.head = NewChunk[T](q.chunkSize)
	q.tail = q.head
}

func (q *Queue[T]) PushFront(value T) {
	q.initialize()

	ok := q.head.PushFront(value)
	if ok {
		return
	}

	c := NewChunk[T](q.chunkSize)
	ok = c.PushFront(value)
	if !ok {
		panic("could not PushFront to a new chunk")
	}

	c.left = q.head
	q.head.right = c
	q.head = c
}

func (q *Queue[T]) PopFront() (T, bool) {
	if q.head == nil {
		var v T
		return v, false
	}

	v, ok := q.head.PopFront()
	if ok {
		return v, ok
	}

	if q.head.left != nil {
		q.head.left.right = nil
	}
	q.head = q.head.left

	if q.head != nil {
		return q.head.PopFront()
	}

	return v, ok
}
