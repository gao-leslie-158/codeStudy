package main
import (
	"fmt"
)

func rob(nums []int) int {
	n := len(nums)
	if n == 1{
		return nums[0]
	}
	dp := make([]int, n)
	dp[0] , dp[1] = nums[0] , max(nums[0],nums[1])
	for i := 2; i < n; i ++{
		dp[i] = max(dp[i-1],dp[i-2]+nums[i])
	}
	return dp[n-1]
}

func max(x , y int) int{
	if x > y{
		return x
	}
	return y
}

func main(){

} 