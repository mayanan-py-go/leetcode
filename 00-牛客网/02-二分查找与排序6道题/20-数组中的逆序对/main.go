package main

import (
	"fmt"
)

/*

*/

func main() {
	array := []int{3, 2, 5, 66, 77, 2000, 1}
	fmt.Println(InversePairs(array))
}
func InversePairs(data []int) int {
	ret := 0
	mergeSort(data, 0, len(data)-1, &ret)
	return ret
}
func mergeSort(li []int, left, right int, ret *int) {
	if left >= right {
		return
	}
	mid := (left+right)/2
	mergeSort(li, left, mid, ret)
	mergeSort(li, mid+1, right, ret)

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
			*ret += mid-i+1
			*ret %= 1e9+7
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
