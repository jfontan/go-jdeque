package jdeque

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChunkFront(t *testing.T) {
	c := newChunk[int](10)

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

func TestChunkBack(t *testing.T) {
	c := newChunk[int](10)

	for i := 0; i < 10; i++ {
		ok := c.PushBack(i)
		require.True(t, ok)
	}

	ok := c.PushBack(10)
	require.False(t, ok)

	for i := 9; i >= 0; i-- {
		v, ok := c.PopBack()
		require.True(t, ok)
		require.Equal(t, i, v)
	}

	v, ok := c.PopBack()
	require.False(t, ok)
	require.Equal(t, 0, v)
}

func TestChunkMixed(t *testing.T) {
	c := newChunk[int](10)

	for i := 0; i < 5; i++ {
		ok := c.PushFront(i)
		require.True(t, ok)
	}

	for i := 0; i < 4; i++ {
		v, ok := c.PopBack()
		require.True(t, ok)
		require.Equal(t, i, v)
	}

	for i := 3; i >= 0; i-- {
		ok := c.PushBack(i)
		require.True(t, ok)
	}

	ok := c.PushBack(-1)
	require.False(t, ok)

	for i := 5; i < 10; i++ {
		ok := c.PushFront(i)
		require.True(t, ok)
	}

	ok = c.PushFront(10)
	require.False(t, ok)

	for i := 0; i < 10; i++ {
		v, ok := c.PopBack()
		require.True(t, ok)
		require.Equal(t, i, v)
	}

	_, ok = c.PopFront()
	require.False(t, ok)
	_, ok = c.PopBack()
	require.False(t, ok)
}
