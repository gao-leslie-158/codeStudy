package main

import (
	"fmt"
)


var (
	path []int
	res [][]int
	used []bool
)

func permute(nums []int) [][]int {
	path , res = make([]int, 0, len(nums)) , make([][]int, 0)
	used = make([]bool, len(nums))
	backtracking(nums,0)
	return res
}

func backtracking(nums []int,cur int){
	// 排列问题需要一个used数组，标记已经选择的元素
	// 而used数组，其实就是记录此时path里都有哪些元素使用了
	// 一个排列里一个元素只能使用一次。
	if cur == len(nums){
		tmp := make([]int, len(path))
		copy(tmp,path)
		res = append(res, tmp)
	}
	for i := 0; i < len(nums); i ++{
		if !used[i] {
			path = append(path, nums[i])
			used[i] = true
			backtracking(nums,cur + 1)
			used[i] = false
			path = path[:len(path) - 1]
		}
	}
}

func main(){

} 