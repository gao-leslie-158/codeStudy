package binarytree

func countNodes(root *TreeNode) int {
	var getNodeNum func(root *TreeNode) int
	getNodeNum = func(root *TreeNode) int {
		if root == nil{
			return 0
		}
		leftNodeNum := getNodeNum(root.Left) // 左
		rightNodeNum := getNodeNum(root.Right) // 右
		return leftNodeNum + rightNodeNum + 1  // 中
	}
	return getNodeNum(root)
}