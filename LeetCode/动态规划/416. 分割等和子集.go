package main

import (
	"fmt"
)

func canPartition(nums []int) bool {
	n := len(nums)
	sum , target := 0 , 0
	for i := 0; i < n; i ++{
		sum += nums[i]
	}
	if sum % 2 != 0{
		return false
	}
	target = sum / 2
	dp := make([]int,target+1)
	for i := 0; i < n; i ++{
		for j := target; j >= nums[i]; j --{
			dp[j] = max(dp[j] , dp[j - nums[i]] + nums[i])
		}
	}
	if dp[target] == target{
		return true
	}
	return false
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



