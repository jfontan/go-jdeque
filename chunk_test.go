package jqueue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChunk(t *testing.T) {
	c := NewChunk[int](10)

	for i := 0; i < 10; i++ {
		ok := c.PushFront(i)
		require.True(t, ok)
	}

	ok := c.PushFront(10)
	require.False(t, ok)

	for i := 9; i >= 0; i-- {
		v, ok := c.PopFront()
		require.True(t, ok)
		require.Equal(t, i, v)
	}

	v, ok := c.PopFront()
	require.False(t, ok)
	require.Equal(t, 0, v)
}
