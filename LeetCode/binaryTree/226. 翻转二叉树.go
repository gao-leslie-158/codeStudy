package binarytree

func invertTree1(root *TreeNode) *TreeNode {
	if root == nil{
		return nil
	}
	root.Left , root.Right = root.Right , root.Left //左右节点交换
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func invertTree2(root *TreeNode) *TreeNode {
	stack := []*TreeNode{}
	node := root
	stack = append(stack, node)
	for len(stack) > 0{
		node = stack[len(stack)-1]	// 根
		stack = stack[:len(stack)-1]
		node.Left , node.Right = node.Right , node.Left  //交换
		if node.Right != nil{
			stack = append(stack, node.Right) // 右
		}
		if node.Left != nil{
			stack = append(stack, node.Left) // 左
		}
	}
	return root
}

// 广度优先遍历，层次遍历
func invertTree(root *TreeNode) *TreeNode {
	if root == nil{
		return root
	}
	queue := []*TreeNode{}
	node := root
	queue = append(queue, node)
	for len(queue) > 0{
		curLen := len(queue) // 当前层长度
		for curLen > 0 {
			node = queue[0]
			queue = queue[1:] // 出队
			curLen -= 1
			node.Left , node.Right = node.Right , node.Left  //交换
			if node.Left != nil{
				queue = append(queue, node.Left)
			}
			if node.Right != nil{
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}
