package binarytree

func levelOrder1(root *TreeNode) [][]int {
	// 借助一个队列，先压入根节点，出队头节点，将出队节点的左右孩子分别压入
	queue := []*TreeNode{}
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue = append(queue, root)
	var tmpArr []int
	for len(queue) > 0 {
		levelLen := len(queue) // 计算本层长度
		for levelLen > 0 {
			node := queue[0]
			queue = queue[1:] // 出队
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			tmpArr = append(tmpArr, node.Val) //存储本层节点值
			levelLen -= 1
		}
		res = append(res, tmpArr)
		tmpArr = []int{} // 清空当前层
	}
	return res
}

func levelOrder(root *TreeNode) [][]int {
	// 递归遍历
	res := [][]int{}
	depth := 0

	var order func(root *TreeNode, depth int)
	order = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if len(res) == depth {
			// res子集存储每一层，跟depth是同步的
			res = append(res, []int{})
		}
		// res[depth]为当前层遍历结果
		res[depth] = append(res[depth], root.Val)
		order(root.Left, depth+1)
		order(root.Right, depth+1)
	}
	order(root, depth)
	return res
}
