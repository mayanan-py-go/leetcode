package main

import (
	"container/list"
	"fmt"
)

/*
给出一个仅包含字符'(',')','{','}','['和']',的字符串，判断给出的字符串是否是合法的括号序列
括号必须以正确的顺序关闭，"()"和"()[]{}"都是合法的括号序列，但"(]"和"([)]"不合法。
*/

func main() {
	fmt.Println(isValid("[](]){}"))
}
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	m1 := map[int32]int32{'{': '}', '[': ']', '(': ')'}
	stack := list.New()
	for _, v := range s {
		if stack.Len() == 0 {
			stack.PushBack(v)
			continue
		}
		tmp := stack.Back().Value.(int32)
		if vv := m1[tmp]; vv == v {
			stack.Remove(stack.Back())
		}else{
			stack.PushBack(v)
		}
	}
	if stack.Len() > 0 {
		return false
	}
	return true
}
















