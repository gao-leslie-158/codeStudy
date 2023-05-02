package main
import (
	"fmt"
)

var (
	path []int
	res [][]int
)

func subsets(nums []int) [][]int {
	path , res = make([]int, 0,len(nums)) , make([][]int, 0)
	backtracking(nums,0)
	return res
}

func backtracking(nums []int,startindex int){
	tmp := make([]int,len(path))
	copy(tmp,path)	// 拷贝当前节点
	fmt.Printf("%d path=%v\n tmp=%v\n",startindex,path,tmp)
	res = append(res, path)	// 收集子集，要放在终止条件的上面，否则会遗漏自己
	if startindex >= len(nums){
		return
	}
	for i := startindex; i < len(nums); i++{
		path = append(path, nums[i])
		backtracking(nums,i+1)
		path = path[:len(path)-1]
	}
}

func main(){
	nums := []int{1,2,3}
	res1 := subsets(nums)
	fmt.Println(res1)
} 