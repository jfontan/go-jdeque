package jqueue

type chunk[T any] struct {
	values      []T
	start       int
	size        int
	left, right *chunk[T]
}

func newChunk[T any](size int) *chunk[T] {
	return &chunk[T]{
		values: make([]T, size),
	}
}

func (c *chunk[T]) canPushFront() bool {
	return (c.start + c.size) < (len(c.values))
}

func (c *chunk[T]) empty() bool {
	return c.size == 0
}

func (c *chunk[T]) pos() int {
	return c.start + c.size
}

func (c *chunk[T]) PushFront(value T) bool {
	if c.empty() && c.start > 0 {
		c.start = 0
	}
	if !c.canPushFront() {
		return false
	}

	c.values[c.pos()] = value
	c.size++
	return true
}

func (c *chunk[T]) PopFront() (T, bool) {
	var v T
	if c.empty() {
		return v, false
	}

	v = c.values[c.pos()-1]
	c.size--
	return v, true
}

func (c *chunk[T]) canPushBack() bool {
	return c.start > 0
}

func (c *chunk[T]) PushBack(value T) bool {
	if c.empty() && c.start < (len(c.values)-1) {
		c.start = len(c.values)
	}
	if !c.canPushBack() {
		return false
	}

	c.start--
	c.size++
	c.values[c.start] = value
	return true
}

func (c *chunk[T]) PopBack() (T, bool) {
	var v T
	if c.empty() {
		return v, false
	}

	v = c.values[c.start]
	c.size--
	c.start++
	return v, true
}
