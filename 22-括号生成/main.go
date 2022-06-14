package main

import "fmt"

/*
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
*/

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	var str string
	var ret []string
	create("", n, n, &str, &ret)
	return ret
}

func create(buffer string, left, right int, str *string, ret *[]string) {
	if left == 0 && right == 0 {
		*str += buffer
		*ret = append(*ret, *str)
		*str = ""
		return
	}
	if left > 0 {
		create(buffer+"(", left-1, right, str, ret)
	}
	if right > left {
		create(buffer+")", left, right-1, str, ret)
	}
}








