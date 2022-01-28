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
BenchmarkJqueue256-16        	       8	 134481116 ns/op	164999233 B/op	  156249 allocs/op
BenchmarkJqueue1024
BenchmarkJqueue1024-16       	       9	 130594665 ns/op	161246101 B/op	   39061 allocs/op
BenchmarkJqueue1048576
BenchmarkJqueue1048576-16    	       8	 127108485 ns/op	157287625 B/op	      37 allocs/op
BenchmarkCarlmjohnson
BenchmarkCarlmjohnson-16     	       5	 253241587 ns/op	587326593 B/op	      12 allocs/op
BenchmarkSekoyo
BenchmarkSekoyo-16           	       6	 193145746 ns/op	357913936 B/op	       4 allocs/op
```