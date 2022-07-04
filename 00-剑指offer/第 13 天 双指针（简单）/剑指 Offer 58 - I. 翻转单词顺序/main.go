package main

import (
	"fmt"
	"strings"
)

/*
输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。为简单起见，标点符号和普通字母一样处理。
例如输入字符串"I am a student. "，则输出"student. a am I"。
https://leetcode.cn/problems/fan-zhuan-dan-ci-shun-xu-lcof/

示例：
输入: "the sky is blue"
输出: "blue is sky the"
 */

func main() {
	fmt.Println(reverseWords("the sky is    blue."))
}
func reverseWords(s string) string {
	/*
	时间复杂度：O(n)
	空间复杂度：O(n)
	 */
	// 删除首位空格
	var ret string
	s = strings.Trim(s, " ")
	i, j := len(s) -1, len(s) - 1
	for i >= 0 {
		for i >= 0 && s[i] != ' '{
			i--
		}
		ret += " " + s[i+1:j+1]
		for i >= 0 && s[i] == ' ' {
			i--
		}
		j = i
	}
	return strings.Trim(ret, " ")
}













