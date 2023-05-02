package binarytree

func minDepth(root *TreeNode) int {
	var getMinDepth func (root *TreeNode) int
	getMinDepth = func(root *TreeNode) int {
		if root == nil{
			return 0
		}
		leftDepth := getMinDepth(root.Left)  // 左
		rightDepth := getMinDepth(root.Right) // 右
											// 中
		// 左子树不为空，右子树为空，不是最低点
		if root.Left != nil && root.Right == nil{
			return leftDepth + 1
		}
		// 左子树为空，右子树不为空，不是最低点
		if root.Left == nil && root.Right != nil{
			return rightDepth + 1
		}
		return 1 + min(leftDepth,rightDepth)
	}
	return getMinDepth(root)
}

func min(x,y int)int{
	if x < y {
		return x
	}
	return y
}