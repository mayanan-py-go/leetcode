package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

示例1：
输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]
*/

func main() {
	arr := []int{4,5,1,6,2,7,3,8}
	fmt.Println(getLeastNumbers(arr, 4))
}
func getLeastNumbers(arr []int, k int) []int {
	shp := New()
	// 初始化堆
	shp.IntSlice = arr
	heap.Init(shp)
	// 堆排序获取前k个元素
	ret := make([]int, k)
	for i, _ := range ret {
		ret[i] = heap.Pop(shp).(int)
	}
	return ret
}
type smallHeap struct {
	sort.IntSlice
}
func New() *smallHeap {
	return new(smallHeap)
}
func (s *smallHeap) Push(x interface{}) {
	s.IntSlice = append(s.IntSlice, x.(int))
}
func (s *smallHeap) Pop() interface{} {
	x := s.IntSlice[len(s.IntSlice)-1]
	s.IntSlice = s.IntSlice[:len(s.IntSlice)-1]
	return x
}
// Less 如果是大根堆需要重写less方法
//func (s *smallHeap) Less(i, j int) bool {
//	return s.IntSlice[i] > s.IntSlice[j]
//}














