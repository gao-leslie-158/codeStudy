package main

import (
	"fmt"
)

func lastStoneWeightII(nums []int) int {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i ++{
		sum += nums[i]
	}
	target := sum / 2
	dp := make([]int,target+1)
	for i := 0; i < n; i ++{
		for j := target; j >= nums[i]; j --{
			//推导公式
			dp[j] = max(dp[j] , dp[j - nums[i]] + nums[i])
		}
	}
	return sum - 2*dp[target]
}

func max(x , y int) int{
	if x > y{
		return x
	}
	return y
}

func main(){
	nums := []int{1,5,11,5}
	res := canPartition(nums)
	fmt.Println(res)
}



