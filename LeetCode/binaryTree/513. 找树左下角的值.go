package binarytree

func findBottomLeftValue1(root *TreeNode) int {
	// 迭代法，用层次遍历，左后一行的第一个元素
	que := []*TreeNode{}
	var ans int
	if root != nil {
		que = append(que, root)
	}
	for len(que) > 0 {
		qLen := len(que)
		for i := 0; i < qLen; i++ {
			node := que[0]
			que = que[1:]
			if i == 0 {
				// 记录最后一行第一个元素
				ans = node.Val
			}
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
	}
	return ans
}

func findBottomLeftValue(root *TreeNode) int {
	// 递归法
	var ans int      // 结果
	var maxDepth int // 最大深度
	var preTravel func(root *TreeNode, depth int)
	preTravel = func(root *TreeNode, depth int) {
		// 递归终止条件
		if root == nil {
			return
		}
		// 记录当前深度最左叶子节点
		if root.Left == nil && root.Right == nil && depth > maxDepth {
			// 因为先遍历左边，所以左边如果有值，右边的同层不会更新结果
			maxDepth = depth
			ans = root.Val
		}
		preTravel(root.Left, depth+1)
		preTravel(root.Right, depth+1)
	}
	preTravel(root, 1)
	return ans
}
