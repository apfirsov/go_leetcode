// https://leetcode.com/problems/kth-largest-element-in-a-stream/

package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	*h = old[0 : len(old)-1]
	return old[len(old)-1]
}

type KthLargest struct {
	minHeap *IntHeap
	k       int
}

func Constructor(k int, nums []int) KthLargest {
	intHeap := IntHeap(nums)
	obj := KthLargest{minHeap: &intHeap, k: k}
	heap.Init(obj.minHeap)
	for len(*obj.minHeap) > k {
		heap.Pop(obj.minHeap)
	}
	return obj
}

func (l *KthLargest) Add(val int) int {
	heap.Push(l.minHeap, val)
	if len(*l.minHeap) > l.k {
		heap.Pop(l.minHeap)
	}
	return (*l.minHeap)[0]
}

func main() {
	k := 3
	nums := []int{4, 5, 8, 2}
	obj := Constructor(k, nums)
	for _, val := range []int{3, 5, 10, 9, 4} {
		fmt.Println(obj.Add(val))
	}
}
