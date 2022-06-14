package main

import "fmt"

func main() {
	ret := myAtoi("-8811954123165465498797")
	fmt.Println(ret)
}

func myAtoi(s string) int {
	var flag int = 1
	var count int
	var length = len(s)

	// 排除前面的空格
	for i := 0; i < length; i++{
		if s[i] == ' '{
			count++
		}else{
			break
		}
	}

	// 全是空格的情况
	if count == length {
		return 0
	}

	// 去完空格之后，如果不是+ - 或0-9则返回0
	if !(s[count] == '+' || s[count] == '-' || (s[count] >= '0' && s[count] <= '9')) {
		return 0
	}

	// 去掉+号或-号，如果是-号，flag设置为-1
	if s[count] == '-' || s[count] == '+'{
		if s[count] == '-'{
			flag = -1
		}
		count++
	}

	// 遍历计算值
	var nums int

	for i := count; i < length; i++{
		if s[i] < '0' || s[i] > '9'{
			break
		}
		nums = nums * 10 + int(s[i] - '0')

		// 超出边界了直接返回
		if flag == 1 && nums > (1<<31)-1{
			nums = (1<<31)-1
			break
		}
		if flag == -1 && nums < -1<<31{
			nums = 1<<31
			break
		}
	}

	// 负数转换
	if flag == -1{
		nums = nums * (-1)
	}

	// 判断是否越界int32
	if nums < -2147483648{
		nums = -2147483648
	}
	if nums > 2147483647{
		nums = 2147483647
	}

	return nums

}
