package src

import (
	"container/heap"
)

type Comparator interface {
	//this>other --> descending order
	//this<other --> ascending order
	Comparator(Comparator) bool
}

type Queue struct {
	queue heap.Interface
}

func New() *Queue {
	pq := &Queue{}
	pq.queue = newHeapMemory()
	return pq
}

func (pq *Queue) Push(qItem Comparator) {
	heap.Push(pq.queue, qItem)
}

func (pq *Queue) Pop() Comparator {

	if pq.queue.Len() <= 0 {
		return nil
	}
	r := heap.Pop(pq.queue)
	return r.(Comparator)
}

type heapMemory struct {
	slice internalSlice
}

func newHeapMemory() *heapMemory {
	return &heapMemory{
		slice: make(internalSlice, 0),
	}
}

type internalSlice []Comparator

func (pq *heapMemory) Len() int { return len(pq.slice) }

func (pq *heapMemory) Less(i, j int) bool {
	return pq.slice[i].Comparator(pq.slice[j])
}

func (pq *heapMemory) Swap(i, j int) {
	pq.slice[i], pq.slice[j] = pq.slice[j], pq.slice[i]
}

func (pq *heapMemory) Push(x interface{}) {
	pq.slice = append(pq.slice, x.(Comparator))
}

func (pq *heapMemory) Pop() interface{} {
	item := pq.slice[len(pq.slice)-1]
	pq.slice = pq.slice[0 : len(pq.slice)-1]
	return item
}
