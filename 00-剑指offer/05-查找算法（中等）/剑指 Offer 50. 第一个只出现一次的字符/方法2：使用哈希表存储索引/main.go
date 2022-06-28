package main

import (
	"fmt"
)

/*
在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。
https://leetcode.cn/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/
*/

func main() {
	s := "abaccdeff"
	fmt.Println(firstUniqChar(s))
}
func firstUniqChar(s string) byte {
	// 使用哈希表存储索引
	var m1 = make(map[int32]int)
	var minIndex = len(s)
	for i, v := range s {
		if _, ok := m1[v]; ok {
			m1[v] = -1
		} else {
			m1[v] = i
		}
	}
	for _, v := range m1 {
		if v == -1 {
			continue
		}
		if v < minIndex {
			minIndex = v
		}
	}
	if minIndex < len(s) {
		return s[minIndex]
	}
	return ' '
}
