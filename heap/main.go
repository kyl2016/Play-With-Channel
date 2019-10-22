package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value string
	score float64
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int{ return len(pq)}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].score < pq[j].score
}

func (pq PriorityQueue) Swap(i, j int){
	pq[i], pq[j] = pq[j], pq[i]
}

func (h *PriorityQueue) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*Item))
}

func (h *PriorityQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	h := &PriorityQueue{
		&Item{
		value: "0",
		score: 0,
	},
		&Item{
			value: "5",
			score: 5,
		},
		&Item{
			value: "2",
			score: 2,
		},
	}
	heap.Init(h)
	heap.Push(h, &Item{
		value:"3",
		score: 3,
	})

	fmt.Printf("minimum: %v\n", (*h)[0])
	//for h.Len() > 0 {
	//	fmt.Printf("%v ", heap.Pop(h))
	//}

	heap.Fix(h, 3)
	for h.Len() > 0 {
		fmt.Printf("%v ", heap.Pop(h))
	}
}
