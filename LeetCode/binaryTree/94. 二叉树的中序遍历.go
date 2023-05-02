package binarytree

func inorderTraversal(root *TreeNode) []int {
	// 利用栈，先将根节点入栈，再将此节点的左节点入栈，直到为叶子结点
	// 弹出栈顶节点，访问，并将其右节点入栈
	st := []*TreeNode{}
	res := make([]int, 0)
	curNode := root
	for curNode != nil || len(st) > 0 {
		if curNode != nil {
			st = append(st, curNode)
			curNode = curNode.Left // 左
		} else {
			curNode = st[len(st)-1]
			st = st[:len(st)-1]
			res = append(res, curNode.Val) // 中
			curNode = curNode.Right        // 右
		}
	}
	return res
}
