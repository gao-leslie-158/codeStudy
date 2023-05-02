package main
import (
	"fmt"
)

func BagProblem(weight []int,value []int,c int) int {
	//初始化
	n := len(weight)
	dp := make([][]int,n)
	for i := 0; i < n; i ++{
		dp[i] = make([]int,c+1)
	}
	//初始化第一行
	for j := 0; j < c; j++{
		if j >= weight[0]{
			dp[0][j] = value[0]		//选取0号物品
		}
	}

	for i := 1; i < n; i ++{
		for j := 0; j <= c; j ++{
			if j < weight[i]{		//装不下
				dp[i][j] = dp[i-1][j]
			}else {
				dp[i][j] = max(dp[i-1][j],dp[i-1][j-weight[i]] + value[i])
			}
		}
	}
	return dp[n-1][c]
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