package test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	pq "github.com/gilomer/priority-queue/priorityqueue"
	"github.com/stretchr/testify/assert"
)

type Item struct {
	a int
	b int
}

func (i Item) Comparator(other pq.Comparator) bool {
	otherToCompare := other.(Item)
	if i.a > otherToCompare.a {
		return true
	} else if i.a < otherToCompare.a {
		return false
	} else if i.b > otherToCompare.b {
		return true
	}
	return false
}

func TestPriorityQueueStruct(t *testing.T) {
	queue := pq.New()
	queue.Push(Item{1, 2})
	queue.Push(Item{2, 2})
	queue.Push(Item{1, 3})
	queue.Push(Item{2, 1})

	assert.Equal(t, Item{2, 2}, queue.Pop())
	assert.Equal(t, Item{2, 1}, queue.Pop())
	assert.Equal(t, Item{1, 3}, queue.Pop())
	assert.Equal(t, Item{1, 2}, queue.Pop())
}

type ObjectWithSimpleAscendingComparator struct {
	msg      string
	priority int
}

func (this *ObjectWithSimpleAscendingComparator) Comparator(other pq.Comparator) bool {
	return this.priority < other.(*ObjectWithSimpleAscendingComparator).priority
}

func TestPriorityQueueAscending(t *testing.T) {
	n := 1000
	pq := pq.New()

	for i := 0; i < n; i++ {
		priority := rand.Intn(100)
		obj := &ObjectWithSimpleAscendingComparator{msg: fmt.Sprintf("Priority %d", priority), priority: priority}
		pq.Push(obj)
	}

	minVal := 0
	i := 0
	for ; i < n; i++ {
		item := pq.Pop()
		cur := item.(*ObjectWithSimpleAscendingComparator).priority
		assert.True(t, cur >= minVal, "Descending sorting failed")
		minVal = cur
	}
	assert.Equal(t, n, i, "Q didn't return the expected amount")
	assert.Nil(t, pq.Pop())
}

type ObjectWithSimpleDescendingComparator struct {
	msg      string
	priority int
}

func (this *ObjectWithSimpleDescendingComparator) Comparator(other pq.Comparator) bool {
	return this.priority > other.(*ObjectWithSimpleDescendingComparator).priority
}

func TestPriorityQueueDescending(t *testing.T) {
	n := 1000
	pq := pq.New()

	for i := 0; i < n; i++ {
		priority := rand.Intn(10)
		obj := &ObjectWithSimpleDescendingComparator{msg: fmt.Sprintf("Priority %d", priority), priority: priority}
		pq.Push(obj)
	}

	maxVal := math.MaxInt64
	i := 0
	for ; i < n; i++ {
		item := pq.Pop()

		cur := item.(*ObjectWithSimpleDescendingComparator).priority
		assert.True(t, cur <= maxVal, "Descending sorting failed")
		maxVal = cur
	}
	assert.Equal(t, n, i, "Q didn't return the expected amount")
	assert.Nil(t, pq.Pop())
}
