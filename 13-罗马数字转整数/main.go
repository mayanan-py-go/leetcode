package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

// 方法二，这个效率要高一些了
func romanToInt(s string) int {
	numsmap := map[byte]int{  // byte类型是uint8
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var res int
	res = numsmap[s[len(s)-1]]

	for i := len(s) - 2; i >= 0; i-- {
		// 核心思想：如何当前值大于等于后面的值，直接相加就行
		curr := numsmap[s[i]]
		if curr >= numsmap[s[i+1]] {
			res += curr
		}else{
			// 如果小于后面的值，那么就减去当前值
			res -= curr
		}
	}

	return res
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
*/

// 方法1，有点慢
/*
执行用时：12 ms, 在所有 Go 提交中击败了34.59%的用户
内存消耗：4.7 MB, 在所有 Go 提交中击败了5.86%的用户
 */
/*
func romanToInt(s string) int {
	numsmap := map[string]int{
		"I": 1,
		"IV": 4,
		"V": 5,
		"IX": 9,
		"X": 10,
		"XL": 40,
		"L": 50,
		"XC": 90,
		"C": 100,
		"CD": 400,
		"D": 500,
		"CM": 900,
		"M": 1000,
	}
	// s := "MCMXCIV" = 1994
	var nums int
	for i := len(s) - 1; i >= 0;{
		if i == 0{
			nums = nums + numsmap[string(s[0])]
			i--
		}else if numsmap[string(s[i])] <= numsmap[string(s[i - 1])]{
			nums = nums + numsmap[string(s[i])]
			i--
		}else{
			nums = nums + numsmap[s[i-1: i + 1]]
			i = i - 2
		}
	}

	return nums
}
*/




