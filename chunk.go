package jqueue

type Chunk[T any] struct {
	values      []T
	start       int
	size        int
	left, right *Chunk[T]
}

func NewChunk[T any](size int) *Chunk[T] {
	return &Chunk[T]{
		values: make([]T, size),
	}
}

func (c *Chunk[T]) canPushFront() bool {
	return (c.start + c.size) < (len(c.values))
}

func (c *Chunk[T]) empty() bool {
	return c.size == 0
}

func (c *Chunk[T]) pos() int {
	return c.start + c.size
}

func (c *Chunk[T]) PushFront(value T) bool {
	if !c.canPushFront() {
		return false
	}
	c.values[c.pos()] = value
	c.size++
	return true
}

func (c *Chunk[T]) PopFront() (T, bool) {
	var v T
	if c.empty() {
		return v, false
	}

	v = c.values[c.pos()-1]
	c.size--
	return v, true
}
