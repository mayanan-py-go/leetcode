package main

import "fmt"

/*
给你一个字符串 s，找到 s 中最长的回文子串。
https://leetcode.cn/problems/longest-palindromic-substring/

示例：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
*/

func main() {
	fmt.Println(longestPalindrome("babad"))
}

func longestPalindrome(s string) string {
	// 中心扩散法
	if len(s) == 0 {
		return ""
	}
	length := 1
	maxLen, maxStart := 0, 0
	for i := 0; i < len(s); i++ {
		left, right := i-1, i+1
		// 向左扩散查找重复值
		for left >= 0 && s[i] == s[left] {
			left--
			length++
		}
		// 向右扩散查找重复值
		for right < len(s) && s[i] == s[right] {
			right++
			length++
		}
		// 两边扩散，查找相同值
		for left >= 0 && right < len(s) && s[left] == s[right] {
			length += 2
			left--
			right++
		}
		if length > maxLen {
			maxLen = length
			maxStart = left
		}
		length = 1
	}
	return s[maxStart+1:maxStart+1+maxLen]
}













