package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	int_max := math.MaxInt
	dp := make([]int, n+1)
	for i,_ := range dp{
		dp[i] = int_max
	}
	dp[0] = 0
	for i := 1; i*i <= n; i ++{
		for j := 1; j <= n; j ++{
			if j >= i*i{
				dp[j] = min(dp[j],dp[j - i*i] + 1)
			}
		}
	}
	if dp[n] == int_max{
		return 0
	}
	return dp[n]
}

func min(x,y int) int{
	if x < y {
		return x
	}
	return y
}

func main(){

}