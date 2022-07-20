package main

import "fmt"

/*

*/

func main() {
	array := []int{3, 2, 5, 66, 77, 2000, 1, 11, 22, 33, 1, 2, 3, 5}
	selectSort(array)
	fmt.Println(array)
}
func selectSort(li []int) {
	// 选择排序
	for i := 0; i < len(li)-1; i++{
		minLoc := i
		for j := i+1; j < len(li); j++ {
			if li[j] < li[minLoc] {
				minLoc = j
			}
		}
		li[i], li[minLoc] = li[minLoc], li[i]
	}
}











