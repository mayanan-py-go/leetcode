package main

import "fmt"

/*

*/

func main() {
	array := []int{3, 2, 1, 5, 66, 77, 2000, 1, 2, 5, 88}
	mergeSort(array, 0, len(array)-1)
	fmt.Println(array)
}
func mergeSortUtil(li []int, low, mid, high int) {
	i := low
	j := mid+1
	var tmp = make([]int, 0)
	for i <= mid && j <= high {
		if li[i] < li[j] {
			tmp = append(tmp, li[i])
			i++
		}else{
			tmp = append(tmp, li[j])
			j++
		}
	}
	// 下面两个循环只会进去一个
	for i <= mid {
		tmp = append(tmp, li[i])
		i++
	}
	for j <= high {
		tmp = append(tmp, li[j])
		j++
	}
	// 将排序好的copy到li中
	copy(li[low:high+1], tmp)
}
func mergeSort(li []int, low, high int) {
	if low < high {
		mid := (low+high) / 2
		mergeSort(li, low, mid)
		mergeSort(li, mid+1, high)
		mergeSortUtil(li, low, mid, high)
	}
}









