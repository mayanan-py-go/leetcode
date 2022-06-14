package main

import "fmt"

func main() {
	b := isValid("{}{[]()}")
	fmt.Println(b)
}

func isValid(s string) bool {
	// 将切片当做栈使用
	s1 := make([]byte, 0, len(s))
	mr := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, tmp := range s{
		if ret, ok := mr[byte(tmp)]; !ok{
			// 如果不是右括号，直接将值加入到切片栈中
			s1 = append(s1, byte(tmp))
		}else{
			// 如果切片中还有值，而且切片栈中的值和对应的左括号一致
			if len(s1) > 0 && s1[len(s1)-1] == ret{
				// 将切片栈的值弹出，继续遍历下一个
				s1 = s1[:len(s1)-1]
			}else{
				// 如果不相等，直接返回
				return false
			}
		}
	}

	// 执行完如果切片栈中还有值，表示未完全匹配
	if len(s1) > 0{
		return false
	}

	return true
}
