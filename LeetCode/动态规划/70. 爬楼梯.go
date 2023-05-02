package main

import (
	"debug/macho"
	"fmt"
)

func climbStairs1(n int) int {
	dp := make([]int,n+1)
	dp[0] , dp[1] = 1 , 2
	for i := 2; i < n; i++{
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}

func climbStairs(n int) int {
	nums := []int{1,2}
	dp := make([]int, n+1)
	dp[0] = 1
	for j := 0; j <= n; j ++{
		for i := 0; i < 2; i ++{
			if j >= nums[i]{
				dp[j] += dp[j - nums[i]]
			}
		}
	}
	return dp[n]
}

func main(){
	res := climbStairs(5)
	fmt.Println(res)
}