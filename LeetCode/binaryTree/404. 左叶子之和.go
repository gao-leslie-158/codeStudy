package binarytree

func sumOfLeftLeaves(root *TreeNode) int {
	// 根节点左叶子和 = 左子树的左叶子和 + 右子树的左叶子和
	// 左叶子：通过父节点判断，左孩子不为空，且左孩子为叶子结点
	var sumLeftValue func(root *TreeNode) int
	sumLeftValue = func(root *TreeNode) int {
		if root == nil{
			return 0
		}
		leftValue := sumLeftValue(root.Left)	// 左
		if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil{
			leftValue += root.Left.Val	// 左叶子和 
		}
		rightValue := sumLeftValue(root.Right)
		return leftValue + rightValue
	}
	return sumLeftValue(root)
}
