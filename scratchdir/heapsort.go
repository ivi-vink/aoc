package main

import (
	"container/heap"
	"fmt"
)

var sortingdata []int = []int{1, 23, 4, 234, 123, 4, 4512, 0}

type myHeap []int

func (h myHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h myHeap) Len() int           { return len(h) }
func (h myHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *myHeap) Pop() any {
    popper := *h
	popped := popper[len(*h)-1]
	*h = popper[:len(*h)-1]
	return popped
}

func (h *myHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func main() {
	orig := myHeap(sortingdata)
	heap.Init(&orig)

	end := orig.Len() - 1
	h := orig
	for h.Len() > 0 {
        heap.Pop(&h)
	}
	fmt.Println(end, orig)
}
