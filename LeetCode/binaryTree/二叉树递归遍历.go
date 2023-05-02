package binarytree


// 前序遍历
func preOrderTraversal(root *TreeNode) (res []int){
	var travelsal func(node *TreeNode)
	travelsal = func (node *TreeNode)  {
		if node == nil{
			return
		}
		res = append(res, node.Val)
		travelsal(node.Left)
		travelsal(node.Right)
	}
	travelsal(root)
	return res
}

// 中序遍历
func midOrderTraversal(root *TreeNode) (res []int){
	var travelsal func(node *TreeNode)
	travelsal = func (node *TreeNode)  {
		if node == nil{
			return
		}
		travelsal(node.Left)
		res = append(res, node.Val)
		travelsal(node.Right)
	}
	travelsal(root)
	return res
}

// 后序遍历
func postOrderTraversal(root *TreeNode) (res []int){
	var travelsal func(node *TreeNode)
	travelsal = func (node *TreeNode)  {
		if node == nil{
			return
		}
		travelsal(node.Left)
		travelsal(node.Right)
		res = append(res, node.Val)
	}
	travelsal(root)
	return res
}