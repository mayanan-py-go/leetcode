package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := isPalindrome(121)
	fmt.Println(b)
}

// 方法3，字符串，字符串前后对调，时间复杂度O(n/2)
func isPalindrome(x int) bool {
	s1 := strconv.Itoa(x)
	runes := []rune(s1)
	for from,to := 0,len(runes)-1; from<to; from,to = from+1,to-1{
		runes[from], runes[to] = runes[to], runes[from]
	}
	if s1 == string(runes){
		return true
	}
	return false
}

// 方法2，字符串，需要申请新的内存，遍历字符串，时间复杂度O(n)
/*
func isPalindrome(x int) bool{
	a := strconv.Itoa(x)
	var ret = make([]rune, 0, len(a))
	for _, v := range a {
		ret = append([]rune{v}, ret...)
	}
	if a == string(ret){
		return true
	}
	return false
}
 */

// 方法1
/*
func isPalindrome(x int) bool {
	// 负数反转一定不是回文数，直接返回
	if x < 0{
		return false
	}

	// 正数反转处理
	var newNum int
	oldNum := x
	for x > 0{
		tmp := x % 10
		x = x / 10
		newNum = newNum * 10 + tmp
	}
	if oldNum == newNum {
		return true
	}
	return false
}
*/
