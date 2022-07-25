package main

/*

*/

func main() {

}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func solve(xianxu []int, zhongxu []int) []int {
	ret := make([]int, 0)
	var dfs func(pre, vin []int, level int)
	dfs = func(pre, vin []int, level int) {
		if len(pre) == 0 {
			return
		}
		if level >= len(ret) {
			ret = append(ret, pre[0])
		}else{
			ret[level] = pre[0]
		}
		// 中序中去查找根节点
		var tmp int
		for i := 0; i < len(vin); i++ {
			if pre[0] == vin[i] {
				tmp = i
				break
			}
		}
		dfs(pre[1:tmp+1], vin[:tmp], level+1)
		dfs(pre[tmp+1:], vin[tmp+1:], level+1)
	}
	dfs(xianxu, zhongxu, 0)
	return ret
}














