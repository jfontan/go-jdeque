package jqueue

type Queue[T any] struct {
	head, tail *chunk[T]
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

	q.head = newChunk[T](q.chunkSize)
	q.tail = q.head
}

func (q *Queue[T]) PushFront(value T) {
	q.initialize()

	ok := q.head.PushFront(value)
	if ok {
		return
	}

	c := newChunk[T](q.chunkSize)
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

func (q *Queue[T]) PushBack(value T) {
	q.initialize()

	ok := q.tail.PushBack(value)
	if ok {
		return
	}

	c := newChunk[T](q.chunkSize)
	ok = c.PushBack(value)
	if !ok {
		panic("could not PushBack to a new chunk")
	}

	c.right = q.tail
	q.tail.left = c
	q.tail = c
}

func (q *Queue[T]) PopBack() (T, bool) {
	if q.tail == nil {
		var v T
		return v, false
	}

	v, ok := q.tail.PopBack()
	if ok {
		return v, ok
	}

	if q.tail.right != nil {
		q.tail.right.left = nil
	}
	q.tail = q.tail.right

	if q.tail != nil {
		return q.tail.PopBack()
	}

	return v, ok
}
