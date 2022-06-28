package main

import "fmt"

/*
在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。
https://leetcode.cn/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/
*/

func main() {
	fmt.Println(firstUniqChar("abbcdca"))
}

type pair struct {
	char  byte
	index int
}

func firstUniqChar(s string) byte {
	n := len(s)
	pos := [26]int{}
	for i := range pos {
		pos[i] = n
	}
	q := make([]pair, 0)
	for i := range s {
		char := s[i] - 'a'
		if pos[char] == n { // 元素第一次出现
			pos[char] = i // 记录第一次出现的索引位置
			q = append(q, pair{char: char, index: i})
		} else { // 元素第二次出现
			pos[char] = n + 1                         // 修改元素的值为需要被移除的值
			for len(q) > 0 && pos[q[0].char] == n+1 { // 当队列长度大于0而且队列中的第一个元素在pos中的值等于n+1时说明重了，需要移除
				q = q[1:] // 队列中弹出一个元素
			}
		}
	}
	// 如果队列长度大于0，队列第一个元素就是第一个只出现一次的字符
	if len(q) > 0 {
		return s[q[0].index]
	}
	return ' '
}
