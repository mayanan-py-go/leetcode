package main

import "fmt"

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
https://leetcode.cn/problems/longest-substring-without-repeating-characters/submissions/
*/

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	var l, r int = 0, 0
	charMap := map[uint8]int{}
	var max int
	var length int = len(s)

	for r = 0; r < length; r++ {
		charMap[s[r]] += 1
		for charMap[s[r]] > 1 {
			charMap[s[l]] -= 1
			l++
		}
		max = Max(max, r-l+1)
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/* 解释
在本题当中，我们可以把一个不包含重复字符的子串当做是原字符串的一个合法区间。比如在字符串 “abcabcbb” 当中[0, 2]就是一个合法区间。
我们用两个记录下标的指针l和r来记录这个区间的左右端点，注意这里的区间我们用的是闭区间。也就是说 l=0，r=2，区间表示好了，怎么移动区间呢？
[a b c] a b c b b
很简单，我们每次往右移动一位，也就是r += 1，区间变成：
[a b c a] b c b b
r往右移动一位，就会读入新的字符，那样整个区间的合法性可能就破坏了。比如我们r加1了之后，读入了a，字符串中多了一个a，那就不是合法区间了。
没关系，我们还有区间的左边界，我们可以再移动区间的左边界，退出一些字符，直到区间重新变成合法区间为止，在这个例子当中，
l只需要移动一位就可以满足条件：
a [b c a] b c b b
也就是说新的区间变成了[1, 3]，这样就完成了区间的移动。如果r移动了之后，依旧没有出现重复字符呢？
没关系，我们继续往下移动就可以了。在这题当中，[0, 0]一定是一个合法的区间，我们可以从[0, 0]开始，
通过移动的方式遍历出所有的合法区间。这些合法区间当中，一定有一个是最终的答案，那么我们的问题也就解决了。
我们再来看一下这种算法的复杂度，它的复杂度是图片。有人会说，我们用了两个指针，不应该也是图片的复杂度吗？
其实不然，看复杂度不能简单只看用了几个循环变量，而需要分析算法当中究竟执行了多少计算量。怎么证明算法复杂度呢？
我们怎么知道窗口到底移动了多少次呢？
不知道移动了多少次也可以，方法很简单，我们分析最坏的情况。算法的起始状态是l=0, r=0。当r=n时算法结束，
我们不知道此时l等于多少，不过没关系。在算法运行的当中，l和r都是递增的，每次窗口移动最多增加1，
那么最多应该执行了2n次（l和r各移动n次）。如此一来，显然这是一个图片的算法。
算法讲完了，还有一个细节没讲清楚，我们怎么维护区间合法呢？
也很简单，我们维护一个map，记录区间内的字符出现了多少次。我们遇到新的字符，就在map中加一，退出字符，就在map中减一。
*/

// 参考文档：https://mp.weixin.qq.com/s?__biz=MzA5NzgzODI5NA%3D%3D&chksm=872bd251b05c5b478e18c5fe9a7ef6915dc7640eabe06c6e78121253a19b4f2181d8a04f9c03&idx=4&mid=2454046696&scene=21&sn=8e42790472ecfce0cd5fc305afb69756#wechat_redirect
