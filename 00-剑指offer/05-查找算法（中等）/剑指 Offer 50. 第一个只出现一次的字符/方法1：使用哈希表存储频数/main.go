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
	// 遍历两次，第一次向字典赋值，第二次遍历的时候，如果遇到等于1的直接返回，也就是第一个字符等于1的
	var m1 = make(map[int32]int)
	for _, b := range s {
		if _, ok := m1[b]; ok {
			m1[b] += 1
		} else {
			m1[b] = 1
		}
	}
	for _, b := range s {
		if v := m1[b]; v == 1 {
			return byte(b)
		}
	}
	return byte(' ')
}
