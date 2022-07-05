package main

/*
给定一个m x n 二维字符网格board 和一个字符串单词word 。如果word 存在于网格中，返回 true ；否则，返回 false 。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
https://leetcode.cn/problems/ju-zhen-zhong-de-lu-jing-lcof/

示例：
输入：board = [
	["A","B","C","E"],
	["S","F","C","S"],
	["A","D","E","E"]
], word = "ABCCED"
输出：true
 */

func main() {
	exist([][]byte{{}}, "abc")
}
func exist(board [][]byte, word string) bool {
	words := []byte(word)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, words, i, j, 0) {
				return true
			}
		}
	}
	return false
}
func dfs(board [][]byte, words []byte, i, j, k int) bool {
	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 || board[i][j] != words[k] {
		return false
	}
	if k == len(words)-1 {
		return true
	}
	//  将 board[i][j] 修改为 空字符 '' ，代表此元素已访问过，防止之后搜索时重复访问。
	board[i][j] = ' '
	res := dfs(board, words, i+1, j, k+1) || dfs(board, words, i-1, j, k+1) ||
		dfs(board, words, i, j+1, k+1) || dfs(board, words, i, j-1, k+1)
	// 还原当前矩阵元素： 将 board[i][j] 元素还原至初始值，即 words[k]
	board[i][j] = words[k]
	return res
}











