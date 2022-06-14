package main

import (
	"fmt"
)

func main() {
	ret := reverse(-150)
	fmt.Println(ret)
}

func reverse(x int) int {
	if x == 0{
		return 0
	}

	// 如果x是负数，先将其转换成正数然后再反转
	isPos := true
	if x < 0{
		isPos = false
		x = 0 - x
	}

	// 定义保存结果的变量
	var ret int
	for x > 0 {
		temp := x % 10
		x = x / 10

		ret = ret * 10 + temp
	}
	// 判断是否超出in32的范围
	//if float64(ret) < math.Pow(-2, 31) || float64(ret) > (math.Pow(2, 31) - 1) {
	//	return 0
	//}
	if ret < -2147483648 || ret > 2147483647{
		return 0
	}

	// 判断原整数是正数还是负数
	if isPos{  // 正数直接返回
		return ret
	}
	// 负数+负号
	return -ret
}
