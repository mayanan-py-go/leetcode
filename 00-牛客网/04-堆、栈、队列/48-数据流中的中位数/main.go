package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。
如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。我们使用Insert()方法读取数据流，
使用GetMedian()方法获取当前读取数据的中位数。
数据范围：数据流中数个数满足 1 \le n \le 1000 \1≤n≤1000  ，大小满足 1 \le val \le 1000 \1≤val≤1000
进阶： 空间复杂度 O(n) \O(n)  ， 时间复杂度 O(nlogn) \O(nlogn)
*/

func main() {
	Insert(15)
	Insert(8)
	Insert(23)
	fmt.Println(GetMedian())
}
// MinHeap 小根堆
type MinHeap struct {
	sort.IntSlice
}
func (m *MinHeap) Push(x interface{}) {
	m.IntSlice = append(m.IntSlice, x.(int))
}
func (m *MinHeap) Pop() interface{} {
	x := m.IntSlice[len(m.IntSlice)-1]
	m.IntSlice = m.IntSlice[:len(m.IntSlice)-1]
	return x
}
// MaxHeap 大根堆
type MaxHeap struct {
	sort.IntSlice
}
func (m *MaxHeap) Push(x interface{}) {
	m.IntSlice = append(m.IntSlice, x.(int))
}
func (m *MaxHeap) Pop() interface{} {
	x := m.IntSlice[len(m.IntSlice)-1]
	m.IntSlice = m.IntSlice[:len(m.IntSlice)-1]
	return x
}
func (m *MaxHeap) Less(i, j int) bool {
	return m.IntSlice[i] > m.IntSlice[j]
}
var minHeap = &MinHeap{}
var maxHeap = &MaxHeap{}
func Insert(num int){
	if minHeap.Len() == maxHeap.Len() {
		if minHeap.Len() == 0 || num >= maxHeap.IntSlice[0] {
			heap.Push(minHeap, num)
		}else{
			heap.Push(maxHeap, num)
			heap.Push(minHeap, heap.Pop(maxHeap))
		}
	}else{
		if num <= minHeap.IntSlice[0] {
			heap.Push(maxHeap, num)
		}else {
			heap.Push(minHeap, num)
			heap.Push(maxHeap, heap.Pop(minHeap))
		}
	}
}
func GetMedian() float64{
	if minHeap.Len() == maxHeap.Len() {
		return float64(minHeap.IntSlice[0]) / 2 + float64(maxHeap.IntSlice[0]) / 2
	}else{
		return float64(minHeap.IntSlice[0])
	}
}

















