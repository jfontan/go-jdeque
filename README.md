# go-jdeque

This queue is implemented with a linked list of chunks. It is a deque and values can be pushed and popped front or back.

Uses generics and needs go >= 1.8.

Example:

```go
q := jdeque.New[int](10)

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
BenchmarkJqueue256-16        	       7	 164305208 ns/op	17978925 B/op	   16746 allocs/op
BenchmarkJqueue1024
BenchmarkJqueue1024-16       	       7	 153615136 ns/op	17352717 B/op	    4189 allocs/op
BenchmarkJqueue1048576
BenchmarkJqueue1048576-16    	       7	 148296433 ns/op	17976419 B/op	       6 allocs/op
BenchmarkCarlmjohnson
BenchmarkCarlmjohnson-16     	       7	 155473335 ns/op	87885730 B/op	       7 allocs/op
BenchmarkSekoyo
BenchmarkSekoyo-16           	       5	 208200644 ns/op	214748377 B/op	       5 allocs/op
```