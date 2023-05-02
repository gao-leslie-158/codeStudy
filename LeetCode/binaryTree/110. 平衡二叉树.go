package binarytree

import "math"

// 给定一个二叉树，判断它是否是高度平衡的二叉树。
func isBalanced(root *TreeNode) bool {
	var getHigh func(root *TreeNode) int
	getHigh = func(root *TreeNode) int {
		// 用-1标记是否平衡
		if root == nil {
			return 0
		}
		leftHigh := getHigh(root.Left) // 左
		if leftHigh == -1 {
			return -1
		}
		rightHigh := getHigh(root.Right) // 右
		if rightHigh == -1 {
			return -1
		}
		if math.Abs(float64(leftHigh-rightHigh)) > 1 {
			return -1
		}
		return 1 + max(leftHigh, rightHigh) // 中
	}
	if getHigh(root) == -1 {
		return false
	}
	return true
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
