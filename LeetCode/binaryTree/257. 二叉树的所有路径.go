package binarytree

import (
	"strconv"
	"strings"
)

var (
	path   []string
	result []string
)

func binaryTreePaths(root *TreeNode) []string {
	//利用回溯，前序遍历思想
	var preTravel func(root *TreeNode)
	preTravel = func(root *TreeNode) {
		// 到叶子结点回溯
		if root.Left == nil && root.Right == nil {
			path = append(path, strconv.Itoa(root.Val)) // 叶子结点也要加入
			tmp := strings.Join(path, "->")
			path = path[:len(path)-1] // 记得出栈
			result = append(result, tmp)
			return
		}
		//前序遍历逻辑
		path = append(path, strconv.Itoa(root.Val)) // 根
		if root.Left != nil {
			preTravel(root.Left) // 左
		}
		if root.Right != nil {
			preTravel(root.Right) // 右
		}
		path = path[:len(path)-1]
	}
	result = make([]string, 0)
	preTravel(root)
	return result
}
