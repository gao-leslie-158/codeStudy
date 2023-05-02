package main
import (
	"fmt"
)

func integerBreak(n int) int {
	dp := make([]int,n+1);
	dp[2] = 1
	for i := 3; i <= n; i++{
		for j := 1; j < i; j++{
			dp[i] = max(dp[i],max(j*(i-j),j*dp[i-j]))
		}
	}
	return dp[n]
}

func max(x,y int) int {
	if x > y{
		return x
	}
	return y
}

func main(){

}