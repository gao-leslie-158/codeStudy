package main
import (
	"fmt"
)

var (
	path []int
	res [][]int
	used []bool
)

func combinationSum2(candidates []int, target int) [][]int {
	path , res = make([]int, 0) , make([][]int, 0)
	used = make([]bool, len(candidates))
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
		// used[i - 1] == true，说明同一树枝candidates[i - 1]使用过
		// used[i - 1] == false，说明同一树层candidates[i - 1]使用过
		if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false{
			continue
		}
		path = append(path, candidates[i])
		used[i] = true
		sum += candidates[i]
		backtracking(candidates,target,sum,i+1) // 这里可重复，不用i+1
		path = path[:len(path)-1]
		sum -= candidates[i]
		used[i] = false
	}
}

func main(){

}