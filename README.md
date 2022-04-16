# Priority Queue
Simple priority queue implementation written in Go

## Example
Use `.new()` to initialize a new priority queue.
The objects used in the queue should be from the same type, and need to implement [Comparator](https://github.com/gilomer/priority-queue/blob/main/priorityqueue/pq.go#L10) inteface.

Use `Push(...)` to add new item to the queue, and `Pop()` to remove items from the queue.
See more examples in the [test file](https://github.com/gilomer/priority-queue/blob/main/test/pq_test.go)

