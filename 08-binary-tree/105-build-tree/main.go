package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	var idx int
	for i, num := range inorder {
		if root.Val == num {
			idx = i
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:idx])+1], inorder[:idx])
	root.Right = buildTree(preorder[len(inorder[:idx])+1:], inorder[idx+1:])

	return root
}

func main() {
	preorder, inorder := []int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}
	buildTree(preorder, inorder)
}
