package main

import "fmt"

/*
有一个长度为 n 的非降序数组，比如[1,2,3,4,5]，将它进行旋转，即把一个数组最开始的若干个元素搬到数组的末尾，
变成一个旋转数组，比如变成了[3,4,5,1,2]，或者[4,5,1,2,3]这样的。请问，给定这样一个旋转数组，求数组中的最小值。
数据范围：1 \le n \le 100001≤n≤10000，数组中任意元素的值: 0 \le val \le 100000≤val≤10000
要求：空间复杂度：O(1)O(1) ，时间复杂度：O(logn)O(logn)
*/

func main() {
	nums := []int{3, 4, 5, 1, 2}
	fmt.Println(minNumberInRotateArray(nums))
}
func minNumberInRotateArray(rotateArray []int) int {
	i, j := 0, len(rotateArray)-1
	for i < j {
		mid := (i+j) / 2
		if rotateArray[mid] > rotateArray[j] {
			i = mid+1
		}else if rotateArray[mid] < rotateArray[j] {
			j = mid
		}else{
			j--
		}
	}
	return rotateArray[i]
}











