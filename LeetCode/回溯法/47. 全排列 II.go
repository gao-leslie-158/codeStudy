package main

import (
	"fmt"
	"sort"
)

var (
	path []int
	res [][]int
	used []bool
)

func permuteUnique(nums []int) [][]int {
	path , res = make([]int, 0, len(nums)) , make([][]int, 0)
	used = make([]bool, len(nums))
	sort.Ints(nums)
	backtracking(nums,0)
	return res
}

func backtracking(nums []int,cur int){
	if cur == len(nums){
		tmp := make([]int, len(path))
		copy(tmp,path)
		res = append(res, tmp)
	}

	for i := 0; i < len(nums); i ++{
		if 1 != 0 && nums[i-1] == nums[i] && used[i-1] == false{
			continue
		}
		if used[i] == false{
			path = append(path, nums[i])
			used[i] = true
			backtracking(nums,cur+1)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
}

func main(){

} 