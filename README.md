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
BenchmarkJqueue256-16        	       6	 176126775 ns/op	165000874 B/op	  156253 allocs/op
BenchmarkJqueue1024
BenchmarkJqueue1024-16       	       6	 172309224 ns/op	161249330 B/op	   39062 allocs/op
BenchmarkJqueue1048576
BenchmarkJqueue1048576-16    	       7	 168460427 ns/op	160583149 B/op	      38 allocs/op
BenchmarkCarlmjohnson
BenchmarkCarlmjohnson-16     	       7	 148484739 ns/op	87885728 B/op	       7 allocs/op
BenchmarkSekoyo
BenchmarkSekoyo-16           	       5	 218807065 ns/op	214748339 B/op	       4 allocs/op
```