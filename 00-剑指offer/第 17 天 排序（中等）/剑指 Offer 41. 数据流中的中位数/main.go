package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。
如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。
https://leetcode.cn/problems/shu-ju-liu-zhong-de-zhong-wei-shu-lcof/

例如：
[2,3,4]的中位数是 3
[2,3] 的中位数是 (2 + 3) / 2 = 2.5
设计一个支持以下两种操作的数据结构：
void addNum(int num) - 从数据流中添加一个整数到数据结构中。
double findMedian() - 返回目前所有元素的中位数。
*/

func main() {
	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(8)
	obj.AddNum(6)
	obj.AddNum(9)
	obj.AddNum(5)
	obj.AddNum(3)
	fmt.Println(obj.FindMedian())
}
type minHeap struct {
	sort.IntSlice
}
type maxHeap struct {
	sort.IntSlice
}
func (m *minHeap) Push(x interface{}) {
	m.IntSlice = append(m.IntSlice, x.(int))
}
func (m *minHeap) Pop() interface{} {
	x := m.IntSlice[len(m.IntSlice) - 1]
	m.IntSlice = m.IntSlice[:len(m.IntSlice)-1]
	return x
}
func (m *maxHeap) Push(x interface{}) {
	m.IntSlice = append(m.IntSlice, x.(int))
}
func (m *maxHeap) Pop() interface{} {
	x := m.IntSlice[len(m.IntSlice) - 1]
	m.IntSlice = m.IntSlice[:len(m.IntSlice)-1]
	return x
}
func (m *maxHeap) Less(i, j int) bool {
	return m.IntSlice[i] > m.IntSlice[j]
}
type MedianFinder struct {
	minHp *minHeap  // 小根堆存储较大的值
	maxHp *maxHeap  // 大根堆存储较小的值
}
func Constructor() MedianFinder {
	return MedianFinder{
		minHp: new(minHeap),
		maxHp: new(maxHeap),
	}
}
func (m *MedianFinder) AddNum(num int) {
	// 添加数字
	if m.minHp.Len() == m.maxHp.Len() {
		if m.minHp.Len() == 0 || num >= m.maxHp.IntSlice[0] {
			heap.Push(m.minHp, num)
		}else {
			heap.Push(m.maxHp, num)
			heap.Push(m.minHp, heap.Pop(m.maxHp))
		}
	}else {
		if num <= m.minHp.IntSlice[0] {
			heap.Push(m.maxHp, num)
		}else {
			heap.Push(m.minHp, num)
			heap.Push(m.maxHp, heap.Pop(m.minHp))
		}
	}
}
func (m *MedianFinder) FindMedian() float64 {
	// 查找中位数
	if m.minHp.Len() == m.maxHp.Len() {
		return float64(m.minHp.IntSlice[0]) / 2 + float64(m.maxHp.IntSlice[0]) / 2
	}else {
		return float64(m.minHp.IntSlice[0])
	}
}
