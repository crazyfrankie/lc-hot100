package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 本题的思路和其他二叉树的题目一致，都是采用递归将问题最小化。
// 从题意来看，是需要找到任意一条路径上的和等于目标值，那么按照递归的思想
// 随着层数越深，当前的总和是在不断增加的，只要有任意一次出现相等，就返回 true，反之 false
// 那么每次递归的操作就是修改当前和（当然官方的做法是每次修改targetSum，意思相近，只是处理不同）
func pathSum(root *TreeNode, targetSum int) bool {
	var cal func(node *TreeNode, current int) bool
	cal = func(node *TreeNode, current int) bool {
		if node == nil {
			return false
		}
		sum := current + node.Val
		if node.Left == nil && node.Right == nil {
			return sum == targetSum
		}
		return cal(node.Left, sum) || cal(node.Right, sum)
	}

	return cal(root, 0)
}

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 8}

	root.Left.Left = &TreeNode{Val: 11}
	root.Right.Left = &TreeNode{Val: 13}
	root.Right.Right = &TreeNode{Val: 4}

	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Left.Left.Right = &TreeNode{Val: 2}
	root.Right.Right.Left = &TreeNode{Val: 1}

	fmt.Println(pathSum(root, 22))
}
