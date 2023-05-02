package main
import (
	"fmt"
)

var (
    path []int  // 存放当前路径的集合
    res [][]int // 存放符合条件结果的集合
)

func combinationSum3(k int, n int) [][]int {
	sum := 0	//用来存放当前路径和
	path , res = make([]int, 0, k) , make([][]int, 0)
	backtrcking(k,n,sum,1)
	return res
}

func backtrcking(k, n int,sum int,start int){
	if len(path) == k{
		if sum == n{
			tmp := make([]int,k)
			copy(tmp,path)
			res = append(res, tmp)
		}
		return
	}

	for i := start; i <= 9; i ++{
		if sum + i > n || 9-i+1 < k - len(path) {	//剪枝
			break
		}
		path = append(path, i)
		sum += i
		backtrcking(k, n, sum, i+1)
		sum -= i
		path = path[:len(path) - 1]
	}
}

func main(){

}