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