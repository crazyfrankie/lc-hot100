package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 本题与 112 不同在于：
// 1. 本题不需要路径从根节点开始，也不需要从叶子节点结束，只要当前和与目标值相等就返回，但路径必须是向下的
// 2. 本题统计的是数量，而不仅仅要求存在

// 那么思路依然是递归，最小化问题

// 将问题转变为，首先计算以任意一个节点为根节点所在的树中的 count：
// 具体来说，不再需要以左右节点均为空作为判断条件，而是只要和相等就增加计数
// 剩下的方式和 112 一致（这里我仍然采用当前和的方式，而没有像官方那样修改当前目标和，我个人觉得当前和比较直观）
// 最后再是递归计算，以该二叉树的每个节点都作为根节点计算总计数，
// 也就是后续的： res += pathSum(root.Left, targetSum) and res += pathSum(root.Right, targetSum)

func pathSum(root *TreeNode, targetSum int) int {
	var res int
	var rootSum func(node *TreeNode, current int) int
	rootSum = func(node *TreeNode, current int) (cnt int) {
		if node == nil {
			return
		}
		sum := current + node.Val
		if sum == targetSum {
			cnt++
		}
		cnt += rootSum(node.Left, sum)
		cnt += rootSum(node.Right, sum)

		return
	}

	if root == nil {
		return 0
	}
	res = rootSum(root, 0)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)

	return res
}

func main() {
	root := &TreeNode{Val: 10}

	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: -3}

	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 11}

	root.Left.Left.Left = &TreeNode{Val: 3}
	root.Left.Left.Right = &TreeNode{Val: -2}
	root.Left.Right.Right = &TreeNode{Val: 1}

	fmt.Println(pathSum(root, 8))
}
