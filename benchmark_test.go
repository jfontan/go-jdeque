package jdeque

import (
	"testing"

	"github.com/carlmjohnson/deque"
	gammazero "github.com/gammazero/deque"
)

func doWork(q *Queue[int]) {
	for i := 0; i < 10_000_000; i++ {
		q.PushFront(i)
	}
	for i := 0; i < 5_000_000; i++ {
		q.PopFront()
	}
	for i := 0; i < 10_000_000; i++ {
		q.PushBack(i)
	}
	for i := 0; i < 5_000_000; i++ {
		q.PopBack()
	}
	for i := 0; i < 10_000_000; i++ {
		q.PopFront()
	}
}

func doWorkDeque(q *deque.Deque[int]) {
	for i := 0; i < 10_000_000; i++ {
		q.PushHead(i)
	}
	for i := 0; i < 5_000_000; i++ {
		q.PopHead()
	}
	for i := 0; i < 10_000_000; i++ {
		q.PushTail(i)
	}
	for i := 0; i < 5_000_000; i++ {
		q.PopTail()
	}
	for i := 0; i < 10_000_000; i++ {
		q.PopHead()
	}
}

func doWorkGammazero(q *gammazero.Deque[int]) {
	for i := 0; i < 10_000_000; i++ {
		q.PushFront(i)
	}
	for i := 0; i < 5_000_000; i++ {
		q.PushFront(i)
	}
	for i := 0; i < 10_000_000; i++ {
		q.PushBack(i)
	}
	for i := 0; i < 5_000_000; i++ {
		q.PopBack()
	}
	for i := 0; i < 10_000_000; i++ {
		q.PopBack()
	}
}

func BenchmarkJqueue256(b *testing.B) {
	q := New[int](256)
	for n := 0; n < b.N; n++ {
		doWork(q)
	}
}

func BenchmarkJqueue1024(b *testing.B) {
	q := New[int](1024)
	for n := 0; n < b.N; n++ {
		doWork(q)
	}
}

func BenchmarkJqueue1048576(b *testing.B) {
	q := New[int](1024 * 1024)
	for n := 0; n < b.N; n++ {
		doWork(q)
	}
}

func BenchmarkCarlmjohnson(b *testing.B) {
	q := deque.Make[int](0)
	for n := 0; n < b.N; n++ {
		doWorkDeque(q)
	}
}

func BenchmarkSekoyo(b *testing.B) {
	q := gammazero.New[int](0)
	for n := 0; n < b.N; n++ {
		doWorkGammazero(q)
	}
}
