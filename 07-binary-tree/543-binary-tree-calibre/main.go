/*
给你一棵二叉树的根节点，返回该树的 直径 。
二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。
两节点之间路径的 长度 由它们之间边数表示。
*/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var dfs func(*TreeNode) int
	res := 0
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		l := dfs(node.Left) + 1
		r := dfs(node.Right) + 1
		res = max(res, l+r)
		return max(l, r)
	}
	dfs(root)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	fmt.Println(diameterOfBinaryTree(root)) // 输出3
}
