package main

import "fmt"

/*

*/

func main() {
	array := []int{3, 2, 1, 5, 66, 77, 2000, 1, 2, 5, 88, 1, 1, 2}
	mergeSort(array, 0, len(array)-1)
	fmt.Println(array)
}
func mergeSort(li []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left+right)/2
	mergeSort(li, left, mid)
	mergeSort(li, mid+1, right)

	i := left
	j := mid+1
	// 设置切片固定长度，避免总是发生扩容
	tmp := make([]int, right-left+1)
	k := 0
	for i <= mid && j <= right {
		if li[i] <= li[j] {
			tmp[k] = li[i]
			i++
			k++
		}else{
			tmp[k] = li[j]
			j++
			k++
		}
	}
	for i <= mid {
		tmp[k] = li[i]
		i++
		k++
	}
	for j <= right {
		tmp[k] = li[j]
		j++
		k++
	}
	copy(li[left:right+1], tmp)
}










