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
	s := "aacabdkacaa"
	fmt.Println(longestPalindrome(s))
}
func longestPalindrome(s string) string {
	var max int
	ss := []byte(s)
	maxByte := make([]byte, 0)
	length := len(ss)
	for i := 0; i < length; i++ {
		for j := i+1; j <= length; j++ {
			tmp := ss[i:j]
			if isPalindromic(tmp) && len(tmp) > max {
				maxByte = tmp
				max = len(tmp)
			}
		}
	}
	return string(maxByte)
}
func isPalindromic(s []byte) bool {
	length := len(s)
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-i-1] {
			return false
		}
	}
	return true
}













