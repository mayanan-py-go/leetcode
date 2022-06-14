package main

import "fmt"

func main() {
	ret := convert("abcdefg", 2)
	fmt.Println(ret)
}

func convert(s string, numRows int) string {
	// 边界条件
	if len(s) <= 2 || numRows == 1{
		return s
	}

	// 初始化结果列表
	res := make([]string, numRows)
	count := 0

	// 计数
	insure := 1

	for _, v := range s{
		res[count] += string(v)
		// 当计数满足转向条件时，转向
		if count == 0 || (count == numRows - 1){
			insure = -insure
		}
		count += -insure
	}

	// 拼凑结果
	result := ""
	for _, v := range res{
		result += v
	}
	return result
}

