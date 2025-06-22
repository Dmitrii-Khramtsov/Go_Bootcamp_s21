package main

import (
	"container/heap"
	"fmt"
	"time"
)

type Present struct {
	Value int
	Size  int
}

type presentHeap []Present

func (h presentHeap) Len() int {
	return len(h)
}

func (h presentHeap) Less(i, j int) bool {
	if h[i].Value > h[j].Value {
		return true
	}

	if h[i].Value == h[j].Value && h[i].Size < h[j].Size {
		return true
	}

	return false
}

func (h presentHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *presentHeap) Push(x any) {
	*h = append(*h, x.(Present))
}

func (h *presentHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	start := time.Now()

	h := &presentHeap{{1500, 1}, {3000, 4}, {2000, 3}, {2000, 1}}
	heap.Init(h)

	heap.Push(h, Present{2500, 2})

	for h.Len() > 0 {
		item := heap.Pop(h).(Present)
		fmt.Printf("Value: %d, Size: %d\n", item.Value, item.Size)
	}

	duration := time.Since(start)
	fmt.Println("Время выполнения:", duration)
}
