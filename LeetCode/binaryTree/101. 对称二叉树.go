package binarytree

func isSymmetric(root *TreeNode) bool {
	/* 递归 */
	var compare func(Left, Right *TreeNode) bool
	compare = func(l, r *TreeNode) bool {
		// 递归终止条件
		if l == nil && r != nil {
			return false
		} else if l != nil && r == nil {
			return false
		} else if l == nil && r == nil {
			return true
		} else if l.Val != r.Val {
			return false
		}
		outside := compare(l.Left, r.Right) // 左：l  右：r
		inside := compare(l.Right, r.Left)  // 左：r  右：l
		isSame := outside && inside         // 左：中  右：中
		return isSame
	}
	if root == nil {
		return true
	}
	return compare(root.Left, root.Right)
}
