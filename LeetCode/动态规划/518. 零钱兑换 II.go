package main
import (
	"fmt"
)
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	n := len(coins)
	dp[0] = 1
	for i := 0; i < n; i ++{	//先物品
		for j := 0; j <= amount; j ++{	//后背包
			if j >= coins[i]{
				dp[j] += dp[j - coins[i]]
			}
		}
		fmt.Println(dp)
	}
	return dp[amount]
}



func main(){
	coins := []int{1, 2, 5}
	amount := 5
	res := change(amount,coins)
	fmt.Println(res)
}