package jdeque

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueueFront(t *testing.T) {
	q := New[int](10)

	for i := 0; i < 100; i++ {
		q.PushFront(i)
	}

	for i := 100 - 1; i >= 0; i-- {
		v, ok := q.PopFront()
		require.True(t, ok)
		require.Equal(t, i, v)
	}

	v, ok := q.PopFront()
	require.False(t, ok)
	require.Equal(t, 0, v)
}

func TestQueueBack(t *testing.T) {
	q := New[int](10)

	for i := 0; i < 100; i++ {
		q.PushBack(i)
	}

	for i := 100 - 1; i >= 0; i-- {
		v, ok := q.PopBack()
		require.True(t, ok)
		require.Equal(t, i, v)
	}

	v, ok := q.PopBack()
	require.False(t, ok)
	require.Equal(t, 0, v)
}

func TestQueueMixed(t *testing.T) {
	q := New[int](10)

	q.PushFront(0)
	q.PushFront(-1)
	v, ok := q.PopBack()
	require.True(t, ok)
	require.Equal(t, 0, v)

	for i := 0; i < 10; i++ {
		q.PushFront(i)
	}

	for i := 0; i < 10; i++ {
		q.PushBack(i)
	}

	var a []int
	for {
		v, ok := q.PopFront()
		if !ok {
			break
		}
		a = append(a, v)
	}

	expected := []int{
		9, 8, 7, 6, 5, 4, 3, 2, 1, 0,
		-1,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	}

	require.Equal(t, expected, a)
}

func TestLen(t *testing.T) {
	q := New[int](10)

	for i := 0; i < 68; i++ {
		q.PushFront(0)
	}

	size := q.Len()
	require.Equal(t, 68, size)
}
