package main

import "fmt"

/*
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数在数组的前半部分，所有偶数在数组的后半部分。
https://leetcode.cn/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/

示例：
输入：nums = [1,2,3,4]
输出：[1,3,2,4]
注：[3,1,2,4] 也是正确的答案之一。
 */

func main() {
	nums := []int{1, 4, 5, 6, 8, 9}
	fmt.Println(exchange(nums))
}
func exchange(nums []int) []int {
	/*
	时间复杂度：O(n)
	空间复杂度：O(1)
	 */
	left, right := 0, len(nums)-1
	for left < right{
		for left < right && nums[left] % 2 == 1{
			left++
		}
		for left < right && nums[right] % 2 == 0 {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	return nums
}
















