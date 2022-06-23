package main

import "fmt"

/*
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
https://leetcode.cn/problems/ti-huan-kong-ge-lcof/
*/

func main() {
	fmt.Println(replaceSpace("abc def 哈哈 555"))
}

func replaceSpace(s string) string {
	var ret string
	for _, ru := range s {
		if ru == ' ' {
			ret += "%20"
		} else {
			ret += string(ru)
		}
	}
	return ret
}
