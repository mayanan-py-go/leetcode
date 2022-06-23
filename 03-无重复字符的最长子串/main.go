package main

import "fmt"

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
https://leetcode.cn/problems/longest-substring-without-repeating-characters/submissions/
*/

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	var l, r int = 0, 0
	charMap := map[uint8]int{}
	var max int
	var length int = len(s)

	for r = 0; r < length; r++ {
		charMap[s[r]] += 1
		for charMap[s[r]] > 1 {
			charMap[s[l]] -= 1
			l++
		}
		max = Max(max, r-l+1)
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
