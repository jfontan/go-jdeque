package jqueue

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
