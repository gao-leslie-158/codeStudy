package main

import (
	"fmt"
	"sort"
)

var (
	path []int
	res [][]int
)

func subsetsWithDup(nums []int) [][]int {
	path ,  res = make([]int, 0,len(nums)) , make([][]int, 0)
	sort.Ints(nums)
	backtracking(nums,0)
	return res
}

func backtracking(nums []int, startindex int){

	tmp := make([]int, len(path))
	copy(tmp,path)
	res = append(res, tmp)

	for i := startindex; i < len(nums); i ++{
		// 同一树层上，不能重复，同一树枝上可以
		if i != startindex && nums[i] == nums[i-1]{
			continue
		}
		path = append(path, nums[i])
		backtracking(nums,i+1)
		path = path[:len(path)-1]
	}
}

func main(){

} 