package jdeque

import "sync"

// Queue implements a deque and values can be pushed or popped from front
// and back.
type Queue[T any] struct {
	head, tail *chunk[T]
	chunkSize  int
	pool       sync.Pool
}

// New creates a new Queue.
func New[T any](chunkSize int) *Queue[T] {
	return &Queue[T]{
		chunkSize: chunkSize,
		pool: sync.Pool{
			New: func() interface{} {
				return newChunk[T](chunkSize)
			},
		},
	}
}

func (q *Queue[T]) initialize() {
	if q.head != nil {
		return
	}

	q.head = q.pool.Get().(*chunk[T])
	q.tail = q.head
}

// PushFront adds a new value to the right of the queue.
func (q *Queue[T]) PushFront(value T) {
	q.initialize()

	ok := q.head.PushFront(value)
	if ok {
		return
	}

	c := q.pool.Get().(*chunk[T])
	ok = c.PushFront(value)
	if !ok {
		panic("could not PushFront to a new chunk")
	}

	c.left = q.head
	q.head.right = c
	q.head = c
}

// PopFront retrieves the rightmost value. Returns false if the queue is empty.
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
	previous := q.head
	q.head = q.head.left

	previous.start = 0
	previous.size = 0
	previous.left = nil
	previous.right = nil
	q.pool.Put(previous)

	if q.head != nil {
		return q.head.PopFront()
	}

	return v, ok
}

// PushBack adds a new value to the left of the queue.
func (q *Queue[T]) PushBack(value T) {
	q.initialize()

	ok := q.tail.PushBack(value)
	if ok {
		return
	}

	c := q.pool.Get().(*chunk[T])
	ok = c.PushBack(value)
	if !ok {
		panic("could not PushBack to a new chunk")
	}

	c.right = q.tail
	q.tail.left = c
	q.tail = c
}

// PopBack retrieves the leftmost value. Returns false if the queue is empty.
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
	previous := q.tail
	q.tail = q.tail.right

	previous.start = 0
	previous.size = 0
	previous.left = nil
	previous.right = nil
	q.pool.Put(previous)

	if q.tail != nil {
		return q.tail.PopBack()
	}

	return v, ok
}

// Len returns the size of the queue.
func (q *Queue[T]) Len() int {
	if q.tail == nil {
		return 0
	}

	var size int
	start := q.tail
	for {
		size += start.size
		if start.right == nil {
			break
		}

		start = start.right
	}

	return size
}
