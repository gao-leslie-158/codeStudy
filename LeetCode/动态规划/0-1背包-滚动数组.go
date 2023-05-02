package main
import (
	"fmt"
)

func BagProblem(weight []int,value []int,c int) int {
	//初始化
	n := len(weight)
	dp := make([]int,c+1)
	for i := 0; i < n; i ++{	//遍历物品
		for j := c; j >= weight[i]; j --{	//倒序遍历背包
			dp[j] = max(dp[j] , dp[j-weight[i]] + value[i])
		}
	}
	return dp[c]
}

func max(x , y int) int {
	if x > y {
		return x
	}
	return y
}


func main(){
	weight := []int{1,2,3,4}
	value := []int{10,15,15,30}
	cap := 7
	maxvalue := BagProblem(weight,value,cap)
	fmt.Println(maxvalue)
}