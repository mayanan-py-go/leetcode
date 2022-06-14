package main

import (
	"fmt"
	"strings"
)

func main() {
	s := intToRoman(1884)
	fmt.Println(s)
}

/*
	字符          数值
	I             1
	V             5
	X             10
	L             50
	C             100
	D             500
	M             1000
1994
*/

func intToRoman(num int) string {
	numsmap := map[int]string{
		1: "I",
		4: "IV",
		5: "V",
		9: "IX",
		10: "X",
		40: "XL",
		50: "L",
		90: "XC",
		100: "C",
		400: "CD",
		500: "D",
		900: "CM",
		1000: "M",
	}
	numsint := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	results := make([]string, 0, 15)
	count := 0

	for i := 0; i < len(numsint) && num != 0; i++{
		count = num / numsint[i]
		num = num - numsint[i] * count
		for count != 0{
			results = append(results, numsmap[numsint[i]])
			count--
		}
	}

	// 转换为字符串
	return strings.Join(results, "")
}







