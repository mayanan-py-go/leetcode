package main

/*
找出数组中重复的数字。
在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，
也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
*/

func main() {

}

func findRepeatNumber(nums []int) int {
	var m1 = make(map[int]byte, 0)
	for _, v := range nums {
		if _, ok := m1[v]; ok {
			return v
		}
		m1[v] = 0
	}
	return 0
}
