package main

import "fmt"

/*
请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。
https://leetcode.cn/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/
输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
*/

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
func lengthOfLongestSubstring(s string) int {
	m1 := make(map[byte]int)
	max := 0
	left, right := 0, 0
	for right = 0; right < len(s); right++ {
		m1[s[right]] += 1
		for m1[s[right]] > 1 {
			m1[s[left]]--
			left++
		}
		max = Max(max, right-left+1)
	}
	return max
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
