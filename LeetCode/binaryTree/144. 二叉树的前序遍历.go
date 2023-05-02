package binarytree

func preorderTraversal(root *TreeNode) []int {
    st := []*TreeNode{}
	res := make([]int, 0)
	if root == nil{
		return res
	}
	st = append(st, root)
	for len(st) > 0{
		node := st[len(st)-1]
		res = append(res, node.Val)
		st = st[:len(st)-1]
		if node.Right != nil{
			st = append(st, node.Right)
		}
		if node.Left != nil{
			st = append(st, node.Left)
		}
	}
	return res
}