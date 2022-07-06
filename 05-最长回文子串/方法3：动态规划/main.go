package main

import "fmt"

/*
给你一个字符串 s，找到 s 中最长的回文子串。
https://leetcode.cn/problems/longest-palindromic-substring/

示例：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
动态规划：如果你发现你有重复调用的过程，动态规划在算完一次之后把答案记录下来，下次再遇到重复的过程直接调用。
动态规划说明了就是空间换时间，将计算结果暂存起来，避免重复计算，作用和工程中redis做缓存有异曲同工之处
动态规划关键是找到初始状态和状态转移方程
初始状态：l=r时，此时dp[l][r]=true
状态转移方程：dp[l][r]=true, 并且l-1和r+1两个位置为相同的字符，此时dp[l-1][r+1]=true
*/

func main() {
	fmt.Println(longestPalindrome("babad"))
}

func longestPalindrome(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}

	// 动态规划存储之前回文串结果, 免得再次查询，节省时间
	var dp = make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
	}

	maxStart, maxLen, maxEnd := 0, 1, 0
	for r := 1; r < length; r++ {
		for l := 0; l < r; l++ {
			if s[l] == s[r] && (r-l<=2 || dp[l+1][r-1]) {
				// 当前l和r所在的子串是回文串
				dp[l][r] = true
				if r-l+1 > maxLen {
					maxLen = r-l+1
					maxStart = l
					maxEnd = r
				}
			}
		}
	}

	return s[maxStart:maxEnd+1]
}













