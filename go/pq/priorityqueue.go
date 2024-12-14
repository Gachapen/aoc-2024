package pq

import (
	"cmp"
	"container/heap"
)

type Item[T any, P cmp.Ordered] struct {
	Value    T
	priority int
	index    int // The index of the item in the heap.
}

type PriorityQueue[T any, P cmp.Ordered] []*Item[T, P]

func (pq PriorityQueue[T, P]) Len() int { return len(pq) }

func (pq PriorityQueue[T, P]) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue[T, P]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue[T, P]) Push(x any) {
	n := len(*pq)
	item := x.(*Item[T, P])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T, P]) Pop() any {
	old := *pq

	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety

	*pq = old[:n-1]

	return item
}

func MakePriorityQueue[T any, P cmp.Ordered]() PriorityQueue[T, P] {
	pq := make(PriorityQueue[T, P], 0)
	heap.Init(&pq)
	return pq
}

func (pq *PriorityQueue[T, P]) PushItem(value T, priority int) {
	heap.Push(pq, &Item[T, P]{Value: value, priority: priority})
}

func (pq *PriorityQueue[T, P]) PopItem() *Item[T, P] {
	return heap.Pop(pq).(*Item[T, P])
}
