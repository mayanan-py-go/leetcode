package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/*

*/

func main() {
	fmt.Println(findKth([]int{1, 8, 6, 8, 5, 3}, 6, 3))
}
func findKth(a []int, n int, K int) int {
	h := &largeHeap{a}
	heap.Init(h)
	for i := 1; i < K; i++ {
		heap.Pop(h)
	}
	return heap.Pop(h).(int)
}
type largeHeap struct {
	sort.IntSlice
}
func (l *largeHeap) Push(x interface{}) {
	l.IntSlice = append(l.IntSlice, x.(int))
}
func (l *largeHeap) Pop() interface{} {
	x := l.IntSlice[len(l.IntSlice)-1]
	l.IntSlice = l.IntSlice[:len(l.IntSlice)-1]
	return x
}
func (l *largeHeap) Less(i, j int) bool {
	return l.IntSlice[i] > l.IntSlice[j]
}











