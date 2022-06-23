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
	var ret = make([]rune, 0, len([]rune(s)))
	for _, tmp := range s {
		if tmp == ' ' {
			ret = append(ret, []rune("%20")...)
		} else {
			ret = append(ret, tmp)
		}
		fmt.Println(tmp)
	}
	return string(ret)
}
