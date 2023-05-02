package binarytree

func postorderTraversal(root *TreeNode) []int {
	// 前序遍历是中左右、调整为中右左，再-逆转一下就是左右中
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
		if node.Left != nil{
			st = append(st, node.Left)
		}
		if node.Right != nil{
			st = append(st, node.Right)
		}
	}
	reverse(res)
	return res
}

func reverse(a []int){
	l , r := 0 , len(a)-1
	for l < r{
		a[l] , a[r] = a[r] , a[l]
		l , r = l+1 , r-1
	}
}
