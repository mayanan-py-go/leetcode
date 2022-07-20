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
	tmp := make([]int, 0)
	for i <= mid && j <= right {
		if li[i] <= li[j] {
			tmp = append(tmp, li[i])
			i++
		}else{
			tmp = append(tmp, li[j])
			j++
		}
	}
	for i <= mid {
		tmp = append(tmp, li[i])
		i++
	}
	for j <= right {
		tmp = append(tmp, li[j])
		j++
	}
	copy(li[left:right+1], tmp)
}










