package main
import (
	"fmt"
)

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n < 2{
		return 0
	}
	dp := make([]int,n+1)
	dp[0] , dp[1] = 0 , 0
	for i :=2; i <= n; i ++{
		dp[i] = min(dp[i-1] + cost[i-1],dp[i-2]+cost[i-2])
	}
	return dp[n]
}

func min(x,y int) int{
	if x < y{
		return x
	}
	return y
}

func main(){

}