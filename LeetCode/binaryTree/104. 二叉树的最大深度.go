package binarytree

func maxDepth1(root *TreeNode) int {
	var getDepth func(node *TreeNode) int
	getDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftDepth := getDepth(node.Left)
		rightDepth := getDepth(node.Right)
		return max(leftDepth, rightDepth) + 1
	}
	return getDepth(root)
}

func maxDepth2(root *TreeNode) (result int) {
	// 前序遍历思想
	var getDepth func(root *TreeNode, depth int)
	getDepth = func(root *TreeNode, depth int) {
		if depth > result { // 中
			result = depth
		}
		if root.Left == nil && root.Right == nil {
			return
		}
		if root.Left != nil { // 左
			getDepth(root.Left, depth+1)
		}
		if root.Right != nil { // 右
			getDepth(root.Right, depth+1)
		}
		return
	}
	getDepth(root, 1)
	return
}
