package main

import "fmt"

func main() {
	s := longestCommonPrefix([]string{"abcd", "abc", "abecd"})
	fmt.Println(s)
}

func longestCommonPrefix(strs []string) string {
	var minLength = 200
	// 获取字符串数组最短长度
	for _, v := range strs{
		if minLength > len(v){
			minLength = len(v)
		}
	}

	// 提取公共部分
	var res = make([]byte, 0)
	for i := 0; i < minLength; i++{
		for ii, v := range strs{
			if ii == 0{
				res = append(res, v[i])
			}else{
				if res[i] == v[i]{
					continue
				} else{
					return string(res[:len(res)-1])
				}
			}
		}
	}

	return string(res)
}
