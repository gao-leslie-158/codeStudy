package main

import (
	"fmt"
	"sort"
)

var (
	path []int
	res [][]int
)

func combinationSum(candidates []int, target int) [][]int {
	path , res = make([]int, 0) , make([][]int, 0)
	sort.Ints(candidates)	// 排序，为剪枝做准备
	backtracking(candidates,target,0,0)
	return res
}

func backtracking(candidates []int, target int, sum int, startindex int){
	if sum == target{
		tmp := make([]int, len(path))
		copy(tmp,path)
		res = append(res, tmp)
		return
	}
	n := len(candidates)
	for i := startindex; i < n; i ++{
		if sum + candidates[i] > target{		// 剪枝
			break
		}
		path = append(path, candidates[i])
		sum += candidates[i]
		backtracking(candidates,target,sum,i) // 这里可重复，不用i+1
		path = path[:len(path)-1]
		sum -= candidates[i]
	}
}

func main(){

}