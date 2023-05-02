package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {
	int_max := math.MaxInt32
	n := len(coins)
	dp := make([]int, amount+1)
	for i , _ := range dp{
		dp[i] = int_max
	}
	dp[0] = 0
	for i := 0; i < n; i ++{
		for j := 0; j <= amount; j ++{
			if j >= coins[i]{
				dp[j] = min(dp[j],dp[j - coins[i]] + 1)
			}
		}
	}
	if dp[amount] == int_max{
		return -1
	}
	return dp[amount]
}

func min(x,y int) int{
	if x < y{
		return x
	}
	return y
}

func main(){

}