# go-jqueue

This queue is implemented with a linked list of chunks. It is a deque and values can be pushed and popped front or back.

Uses generics and needs go >= 1.8.

Example:

```go
q := jqueue.New[int](10)

q.PushFront(0)
q.PushBack(1)

v, ok := q.PopFront() // 0, true
v, ok = q.PopFront() // 1, true
v, ok = q.PopFront() // 0, false
```

## Benchmarks

This is comparing different chunk sizes and the following libraries:

* https://github.com/carlmjohnson/deque
* https://github.com/sekoyo/deque

```
BenchmarkJqueue256
BenchmarkJqueue256-16        	       6	 168462951 ns/op	164999348 B/op	  156250 allocs/op
BenchmarkJqueue1024
BenchmarkJqueue1024-16       	       9	 133794490 ns/op	161246101 B/op	   39061 allocs/op
BenchmarkJqueue1000000
BenchmarkJqueue1000000-16    	       9	 118682889 ns/op	156588565 B/op	      37 allocs/op
BenchmarkDeque
BenchmarkDeque-16            	       4	 302940608 ns/op	587282408 B/op	      16 allocs/op
BenchmarkGammazero
BenchmarkGammazero-16        	       6	 241483642 ns/op	357913936 B/op	       4 allocs/op
```