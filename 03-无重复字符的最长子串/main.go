package main

import (
	"fmt"
)

func main() {
	ret := lengthOfLongestSubstring("abacbded")
	fmt.Println(ret)
}

func lengthOfLongestSubstring(s string) int {
	n, l, p := 0, 0, -1
	substr := make([]rune, 0, len(s))
	for _, v := range s{

		p = func(runes []rune, r rune) int {
			for i, vv := range runes{
				if vv == r{
					return i
				}
			}
			return -1
		}(substr, v)

		if p > -1{
			l = len(substr)
			if l > n{
				n = l
			}
			substr = append(substr[p+1:], v)
		}else{
			substr = append(substr, v)
		}
	}
	l = len(substr)
	if l > n{
		n = l
	}
	return n
}

