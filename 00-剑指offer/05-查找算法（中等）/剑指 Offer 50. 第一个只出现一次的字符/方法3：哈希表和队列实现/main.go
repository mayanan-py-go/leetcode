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
	char  int32
	index int
}

func firstUniqChar(s string) byte {
	// hashmap+队列的方式实现
	m1 := make(map[int32]int)
	q := make([]pair, 0)
	for i, v := range s {
		if _, ok := m1[v]; !ok {
			m1[v] = i
			q = append(q, pair{char: v, index: i})
		} else {
			m1[v] = -1
			for len(q) > 0 && m1[q[0].char] == -1 {
				q = q[1:]
			}
		}
	}
	if len(q) > 0 {
		return s[q[0].index]
	}
	return ' '
}
